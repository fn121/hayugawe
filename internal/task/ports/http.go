package ports

import (
	"encoding/json"
	"net/http"

	"github.com/fn121/hayugawe/internal/task/app"
	"github.com/fn121/hayugawe/internal/task/app/command"
	"github.com/fn121/hayugawe/internal/task/app/query"
	"github.com/fn121/hayugawe/internal/task/domain/task"
)

func NewHttpServer(app app.Application) HttpServer{
	return HttpServer{app: app}
}

type HttpServer struct {
	app app.Application
}

func (h *HttpServer) RegisterRoutes(mux *http.ServeMux) {

	mux.HandleFunc("GET /api/tasks", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		tasks, err := h.app.Queries.GetTaskList.Handle(r.Context(), query.GetTaskList{})
		
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

		err := h.app.Commands.CreateTask.Handle(r.Context(), command.CreateTask{TaskData: taskData})
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

		err := h.app.Commands.MarkTaskAsDone.Handle(r.Context(), command.MarkTaskAsDone{
			TaskID: taskId,
		})
		if err != nil {
			// return task not found
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(HttpError{Message: err.Error()})
		}
		
		w.WriteHeader(http.StatusNoContent)
	})

	mux.HandleFunc("PATCH /api/tasks/{taskId}/status/todo", func(w http.ResponseWriter, r *http.Request) {
		taskId := r.PathValue("taskId");
		
		err := h.app.Commands.MarkTaskAsTodo.Handle(r.Context(), command.MarkTaskAsTodo{TaskID: taskId})
		if err != nil {
			// return task not found
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(HttpError{Message: err.Error()})
		}
		

		w.WriteHeader(http.StatusNoContent)
	})

	// Test CICD. Auto build, and deploy when push to main branch. Safe to be deleted
	mux.HandleFunc("GET /api/testcicd", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(`{"message": "build and deploy pipeline success"}`)
	})
}
