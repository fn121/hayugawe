package app

import (
	"context"

	"github.com/fn121/hayugawe/internal/task/domain/task"
)

type TaskService struct {
	repo task.Repository
}

func NewTaskService(repo task.Repository) TaskService {
	return TaskService{repo}
}

func (ts *TaskService) GetTask(taskID string) (*task.Task, error) {
	ctx := context.Background()
	return ts.repo.FindByID(ctx, taskID)
}

func (ts *TaskService) GetTaskList() ([]task.Task, error) {
	ctx := context.Background()
	tasks, err := ts.repo.Find(ctx)

	return tasks, err
}