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

func RunHTTPServer(handler http.Handler) {
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
	repository := adapters.NewMemoryTaskRepositoryWithSample()
	app := newApplication(repository, repository)
	apiMux := http.NewServeMux()
	mainMux := http.NewServeMux()

	handler := ports.HandlerFromMux(ports.NewHttpServer(app), apiMux)
	mainMux.Handle("/api/", http.StripPrefix("/api", handler))

	RunHTTPServer(mainMux)
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