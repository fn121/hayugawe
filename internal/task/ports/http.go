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

// GetTaskList Get task list
// (GET /tasks)
func (h HttpServer) GetTaskList(w http.ResponseWriter, r *http.Request) {
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
}

// CreateTask Create a task
// (POST /tasks)
func (h HttpServer) CreateTask(w http.ResponseWriter, r *http.Request) {
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
}

// MarkTaskAsDone Mark task as done
// (PATCH /tasks/{taskId}/status/done)
func (h HttpServer) MarkTaskAsDone(w http.ResponseWriter, r *http.Request, taskId string){

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
}

// MarkTaskAsTodo Mark task as todo
// (PATCH /tasks/{taskId}/status/todo)
func (h HttpServer) MarkTaskAsTodo(w http.ResponseWriter, r *http.Request, taskId string){
	err := h.app.Commands.MarkTaskAsTodo.Handle(r.Context(), command.MarkTaskAsTodo{TaskID: taskId})
	if err != nil {
		// return task not found
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(HttpError{Message: err.Error()})
	}
	

	w.WriteHeader(http.StatusNoContent)
}

// TestCICD Test CI/CD deployment
// (GET /testcicd)
func (h HttpServer) TestCICD(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(`{"message": "build and deploy pipeline success"}`)
}