package ports

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/fn121/hayugawe/internal/task/app"
	"github.com/fn121/hayugawe/internal/task/domain/task"
)

type TaskService struct {
	list []task.Task
}

func (ts *TaskService) FindByID(id string) (*task.Task, error) {
	for idx, t := range ts.list {
		if t.ID == id {
			return &ts.list[idx], nil
		}
	}
	return nil, fmt.Errorf("task with id %q not found", id)
}

func NewHttpServer(service app.TaskService) HttpServer{
	return HttpServer{}
}

type HttpServer struct {
	service TaskService
}

type HttpError struct {
	Message string `json:"message"`
}

func (h *HttpServer) RegisterRoutes(mux *http.ServeMux) {

	mux.HandleFunc("GET /api/tasks", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(h.service.list)
	})

	mux.HandleFunc("POST /api/tasks", func(w http.ResponseWriter, r *http.Request) {
		var taskData task.PostTask
		if readBodyErr := json.NewDecoder(r.Body).Decode(&taskData); readBodyErr != nil {
			w.Header().Set("Content-Type", "application/json")
			// send 400 bad request
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(HttpError{Message: "Bad request"})
		}

		newTask := task.NewTask(taskData)
		h.service.list = append(h.service.list, *newTask )

		w.WriteHeader(http.StatusNoContent)
	})

	mux.HandleFunc("PATCH /api/tasks/{taskId}/status/done", func(w http.ResponseWriter, r *http.Request) {
		taskId := r.PathValue("taskId");
		
		foundTask, err := h.service.FindByID(taskId)
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
		
		foundTask, err := h.service.FindByID(taskId)
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
