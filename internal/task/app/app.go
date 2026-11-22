package app

import (
	"github.com/fn121/hayugawe/internal/task/app/command"
	"github.com/fn121/hayugawe/internal/task/app/query"
)

type Application struct {
	Commands Commands
	Queries Queries
}

type Commands struct {
	MarkTaskAsDone command.MarkTaskAsDoneHandler
	MarkTaskAsTodo command.MarkTaskAsTodoHandler
	CreateTask command.CreateTaskHandler
}

type Queries struct {
	GetTaskList query.GetTaskListHandler
}