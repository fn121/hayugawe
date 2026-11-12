package main

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

type Task struct {
	ID string `json:"id"`
	Title string `json:"title"`
}

type TaskService struct {
	list []Task
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
	}	
}

func main() {
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

	http.ListenAndServe(":8080", mux)
}