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

func TestTaskRepository_UpdateTask(t *testing.T) {
	task := newTaskTodo()
	r := adapters.NewMemoryTaskRepository(adapters.TaskList{adapters.TaskModel(*task)})
	ctx := context.Background()


	// update
	task.MarkDone()
	errUpdate := r.UpdateTask(ctx, task.ID, *task)
	require.NoError(t, errUpdate)

	// assert
	updatedTask, err := r.FindTaskByID(ctx, task.ID)
	require.NoError(t, err)
	assert.Equal(t, true, updatedTask.IsDone)
}

func TestTaskRepository_UpdateTask_not_exist_return_error(t *testing.T) {
	newTask := newTaskTodo()
	r := adapters.NewMemoryTaskRepository(adapters.TaskList{})
	ctx := context.Background()

	errUpdate := r.UpdateTask(ctx, newTask.ID, *newTask)
	require.ErrorIs(t, errUpdate, task.ErrTaskNotFound)
}

func newTask() *task.Task {
	return task.MustNewTask(task.PostTask{Title: "Do home work"})
}

func newTaskTodo() *task.Task {
	return &task.Task{ID: "should-uuid", Title: "task todo", IsDone: false}
}

func newTaskDone() *task.Task {
	return &task.Task{ID: "should-uuid", Title: "task todo", IsDone: true}
}


func assertCountTasks(t *testing.T, r task.Repository, wantCount int) {
	t.Helper()
	ctx := context.Background()
	list, err := r.FindTask(ctx)
	assert.NoError(t, err)

	assert.Equal(t, len(list), wantCount)
}