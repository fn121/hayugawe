package main

import (
	"log"
	"net/http"
	"os"

	"github.com/fn121/hayugawe/internal/task/adapters"
	"github.com/fn121/hayugawe/internal/task/app"
	"github.com/fn121/hayugawe/internal/task/app/command"
	"github.com/fn121/hayugawe/internal/task/app/query"
	"github.com/fn121/hayugawe/internal/task/domain/task"
	"github.com/fn121/hayugawe/internal/task/ports"
)

func RunHTTPServer(handler *http.ServeMux) {
	log.Print("starting server...")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting port to %s", port)
	}

	log.Printf("listening on port %s", port)
	log.Panic(http.ListenAndServe(":"+port, handler))
}

func main() {
	handler := http.NewServeMux()

	repository := adapters.NewMemoryTaskRepositoryWithSample()
	app := newApplication(repository, repository)
	
	httpServer := ports.NewHttpServer(app)
	httpServer.RegisterRoutes(handler)

	RunHTTPServer(handler)
}

func newApplication(repo task.Repository, queryRepo query.GetTaskListReadModel ) app.Application {

	return app.Application{
		Commands: app.Commands{
			MarkTaskAsDone: command.NewMarkTaskAsDoneHandler(repo),
			MarkTaskAsTodo: command.NewMarkTaskAsTodoHandler(repo),
			CreateTask: command.NewCreateTaskHandler(repo),
		},
		Queries: app.Queries{
			GetTaskList: query.NewGetTaskListHandler(queryRepo),
		},
	}
}