package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"
)

type Task struct {
	ID string `json:"id"`
	Title string `json:"title"`
	IsDone bool `json:"is_done"`
}

func(t *Task) MarkDone() {
	t.IsDone = true
}


func(t *Task) MarkTodo() {
	t.IsDone = false
}


func(t *Task) ToggleStatus() {
	t.IsDone = !t.IsDone
}


type TaskService struct {
	list []Task
}

func (ts *TaskService) FindByID(id string) (*Task, error) {
	for idx, t := range ts.list {
		if t.ID == id {
			return &ts.list[idx], nil
		}
	}
	return nil, fmt.Errorf("task with id %q not found", id)
}

type PostTask struct {
	Title string `json:"title"`
}

type HttpError struct {
	Message string `json:"message"`
}

func NewTask(taskData PostTask) *Task{
	return &Task{
		ID: uuid.NewString(),
		Title: taskData.Title,
		IsDone: false,
	}	
}

func main() {
	log.Print("starting server...")
	mux := http.NewServeMux()

	service := TaskService{list: []Task{{ID: uuid.New().String(), Title: "Do something" }}}
	mux.HandleFunc("GET /api/tasks", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(service.list)
	})

	mux.HandleFunc("POST /api/tasks", func(w http.ResponseWriter, r *http.Request) {
		var taskData PostTask
		if readBodyErr := json.NewDecoder(r.Body).Decode(&taskData); readBodyErr != nil {
			w.Header().Set("Content-Type", "application/json")
			// send 400 bad request
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(HttpError{Message: "Bad request"})
		}

		newTask := NewTask(taskData)
		service.list = append(service.list, *newTask )

		w.WriteHeader(http.StatusNoContent)
	})

	mux.HandleFunc("PATCH /api/tasks/{taskId}/status/done", func(w http.ResponseWriter, r *http.Request) {
		taskId := r.PathValue("taskId");
		
		foundTask, err := service.FindByID(taskId)
		if err != nil {
			// return task not found
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(HttpError{Message: err.Error()})
		}
		
		foundTask.MarkDone()

		w.WriteHeader(http.StatusNoContent)
	})

	mux.HandleFunc("PATCH /api/tasks/{taskId}/status/todo", func(w http.ResponseWriter, r *http.Request) {
		taskId := r.PathValue("taskId");
		
		foundTask, err := service.FindByID(taskId)
		if err != nil {
			// return task not found
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(HttpError{Message: err.Error()})
		}
		
		foundTask.MarkTodo()

		w.WriteHeader(http.StatusNoContent)
	})

	// Test CICD. Auto build, and deploy when push to main branch. Safe to be deleted
	mux.HandleFunc("GET /api/testcicd", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(`{"message": "build and deploy pipeline success"}`)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting port to %s", port)
	}

	log.Printf("listening on port %s", port)
	http.ListenAndServe(":"+port, mux)
}