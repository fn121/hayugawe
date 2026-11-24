package command

import (
	"context"
	"fmt"

	"github.com/fn121/hayugawe/internal/task/domain/task"
)

type CreateTask struct {
	TaskData task.PostTask
}

type CreateTaskHandler struct {
	repo task.Repository
}

func NewCreateTaskHandler(repo task.Repository) CreateTaskHandler {
	if repo == nil {
		panic("nil repo")
	}

	return CreateTaskHandler{repo}
}

func (h CreateTaskHandler) Handle(ctx context.Context, cmd CreateTask) (taskID string, error error) {
	
	newTask, err := task.NewTask(cmd.TaskData)
	if err != nil {
		return "", fmt.Errorf("fail to create task: %w", err)
	}

	errSave := h.repo.AddTask(ctx, *newTask)
	if errSave != nil {
		return "", fmt.Errorf("fail to save: %w", errSave)
	}
	return newTask.ID, nil
}
