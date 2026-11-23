package adapters_test

import (
	"context"
	"testing"

	"github.com/fn121/hayugawe/internal/task/adapters"
	"github.com/fn121/hayugawe/internal/task/domain/task"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTaskRepository_AddTask(t *testing.T) {
	r := adapters.NewMemoryTaskRepository(adapters.TaskList{})
	// pre assumptions
	assertCountTasks(t, r, 0)

	ctx := context.Background()
	err := r.AddTask(ctx, *newTask())
	require.NoError(t, err)

	assertCountTasks(t, r, 1)
}

func newTask() *task.Task {
	return task.MustNewTask(task.PostTask{Title: "Do home work"})
}

func assertCountTasks(t *testing.T, r task.Repository, wantCount int) {
	t.Helper()
	ctx := context.Background()
	list, err := r.FindTask(ctx)
	assert.NoError(t, err)

	assert.Equal(t, len(list), wantCount)
}