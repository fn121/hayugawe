package task

import "context"

type Repository interface {
	AddTask(ctx context.Context, task Task) error
	UpdateTask(ctx context.Context, taskID string, task Task) error
	FindTaskByID(ctx context.Context, taskID string) (*Task, error)
	FindTask(ctx context.Context) ([]Task, error)
}