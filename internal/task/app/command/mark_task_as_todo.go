package command

import (
	"context"
	"fmt"

	"github.com/fn121/hayugawe/internal/task/domain/task"
)

type MarkTaskAsTodo struct {
	TaskID string
}

type MarkTaskAsTodoHandler struct {
	repo task.Repository
}

func NewMarkTaskAsTodoHandler(repo task.Repository) MarkTaskAsTodoHandler {
	if repo == nil {
		panic("nil repo")
	}

	return MarkTaskAsTodoHandler{repo}
}

func (h MarkTaskAsTodoHandler) Handle(ctx context.Context, cmd MarkTaskAsTodo) (error) {
	task, err := h.repo.FindByID(ctx, cmd.TaskID)
	if err != nil {
		return fmt.Errorf("failed to find task: %w", err)
	}

	task.MarkTodo()

	return h.repo.UpdateByID(
		ctx,
		cmd.TaskID,
		*task,
	)
}
