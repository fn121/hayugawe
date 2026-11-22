package query

import (
	"context"
)

// define read model interface
type GetTaskListReadModel interface {
	GetTaskList() ([]Task, error)
}

func NewGetTaskListHandler(readModel GetTaskListReadModel) GetTaskListHandler {
	return GetTaskListHandler{readModel}
}

// kind of required data for the query
type GetTaskList struct {}

type GetTaskListHandler struct {
	readModel GetTaskListReadModel
}

func (h GetTaskListHandler) Handle(ctx context.Context, query GetTaskList ) ([]Task, error) {
	return h.readModel.GetTaskList()
}