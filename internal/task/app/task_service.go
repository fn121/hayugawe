package app

import (
	"fmt"

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