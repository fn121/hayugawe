package service

import (
	"github.com/fn121/hayugawe/internal/task/adapters"
	"github.com/fn121/hayugawe/internal/task/app"
	"github.com/fn121/hayugawe/internal/task/app/command"
	"github.com/fn121/hayugawe/internal/task/app/query"
	"github.com/fn121/hayugawe/internal/task/domain/task"
)

func NewApplication() app.Application {
	taskRepo := adapters.NewMemoryTaskRepositoryWithSample()
	return newApplication(taskRepo, taskRepo)
}

func NewComponentTestApplication() app.Application {
	taskRepo := adapters.NewMemoryTaskRepositoryWithSample()
	return newApplication(taskRepo, taskRepo)
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