package task

import "context"

type Repository interface {
	Save(ctx context.Context, task Task) error
	UpdateByID(ctx context.Context, taskID string, task Task) error
	FindByID(ctx context.Context, taskID string) (*Task, error)
	Find(ctx context.Context) ([]Task, error)
}