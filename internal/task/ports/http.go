package ports

import (
	"encoding/json"
	"net/http"

	"github.com/fn121/hayugawe/internal/task/app"
	"github.com/fn121/hayugawe/internal/task/domain/task"
)

func NewHttpServer(service app.TaskService) HttpServer{
	return HttpServer{taskService: service}
}

type HttpServer struct {
	taskService app.TaskService
}

type HttpError struct {
	Message string `json:"message"`
}

func (h *HttpServer) RegisterRoutes(mux *http.ServeMux) {

	mux.HandleFunc("GET /api/tasks", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		tasks, err := h.taskService.GetTaskList()
		
		if err != nil {
			// w.Header().Set("Content-Type", "application/json")
			// send 400 bad request
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(HttpError{Message: "Bad request"})
			return
		}

		json.NewEncoder(w).Encode(tasks)
	})

	mux.HandleFunc("POST /api/tasks", func(w http.ResponseWriter, r *http.Request) {
		var taskData task.PostTask
		if readBodyErr := json.NewDecoder(r.Body).Decode(&taskData); readBodyErr != nil {
			w.Header().Set("Content-Type", "application/json")
			// send 400 bad request
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(HttpError{Message: "Bad request"})
		}

		err := h.taskService.CreateTask(taskData)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			// send 400 bad request
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(HttpError{Message: "Bad request"})
		}


		w.WriteHeader(http.StatusNoContent)
	})

	mux.HandleFunc("PATCH /api/tasks/{taskId}/status/done", func(w http.ResponseWriter, r *http.Request) {
		taskId := r.PathValue("taskId");
		
		foundTask, err := h.taskService.GetTask(taskId)
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
		
		foundTask, err := h.taskService.GetTask(taskId)
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
}
