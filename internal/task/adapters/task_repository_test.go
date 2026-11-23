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
	t.Parallel()
	r := adapters.NewMemoryTaskRepository(adapters.TaskList{})

	newTask := newTask()
	ctx := context.Background()
	err := r.AddTask(ctx, *newTask)
	require.NoError(t, err)

	assertTaskInRepository(t, r, newTask)
}

func TestTaskRepository_UpdateTask(t *testing.T) {
	t.Parallel()
	task := newTask()
	task.MarkTodo()
	r := adapters.NewMemoryTaskRepository(adapters.TaskList{adapters.TaskModel(*task)})
	ctx := context.Background()


	// update
	task.MarkDone()
	errUpdate := r.UpdateTask(ctx, task.ID, *task)
	require.NoError(t, errUpdate)

	// assert
	assertTaskInRepository(t, r, task)
}

func TestTaskRepository_UpdateTask_not_exist_return_error(t *testing.T) {
	t.Parallel()
	newTask := newTask()
	r := adapters.NewMemoryTaskRepository(adapters.TaskList{})
	ctx := context.Background()

	errUpdate := r.UpdateTask(ctx, newTask.ID, *newTask)
	assert.ErrorIs(t, errUpdate, task.ErrTaskNotFound)
}

func newTask() *task.Task {
	return task.MustNewTask(task.PostTask{Title: "Do home work"})
}


func assertTaskInRepository(t *testing.T, repo task.Repository, task *task.Task) {
	require.NotNil(t, task)

	ctx := context.Background()
	taskFromRepo, err := repo.FindTaskByID(ctx, task.ID)
	require.NoError(t, err)

	assert.Equal(t, task, taskFromRepo)
}
