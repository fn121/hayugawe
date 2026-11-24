package command

import (
	"context"
	"fmt"

	"github.com/fn121/hayugawe/internal/task/domain/task"
)

type MarkTaskAsDone struct {
	TaskID string
}

type MarkTaskAsDoneHandler struct {
	repo task.Repository
}

func NewMarkTaskAsDoneHandler(repo task.Repository) MarkTaskAsDoneHandler {
	if repo == nil {
		panic("nil repo")
	}

	return MarkTaskAsDoneHandler{repo}
}

func (h MarkTaskAsDoneHandler) Handle(ctx context.Context, cmd MarkTaskAsDone) (error) {
	task, err := h.repo.FindTaskByID(ctx, cmd.TaskID)
	if err != nil {
		return fmt.Errorf("failed to find task: %w", err)
	}

	task.MarkDone()

	return h.repo.UpdateTask(
		ctx,
		cmd.TaskID,
		*task,
	)
}
