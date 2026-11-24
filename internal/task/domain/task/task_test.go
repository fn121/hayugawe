package task_test

import (
	"testing"

	"github.com/fn121/hayugawe/internal/task/domain/task"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewTask(t *testing.T) {
	taskData := task.PostTask{
		Title: "Add feature: task project",
	}

	task, err := task.NewTask(taskData)
	require.NoError(t, err)

	assert.Equal(t, taskData.Title, task.Title)
}

func TestNewTask_invalid(t *testing.T) {
	taskData := task.PostTask{
		Title: "",
	}

	_, err := task.NewTask(taskData)
	require.Error(t, err)
}

func TestNewTask_MarkDone(t *testing.T) {
	task := newExampleTaskTodo()

	task.MarkDone()

	assert.True(t, task.IsDone)
}

func TestNewTask_MarkTodo(t *testing.T) {
	task := newExampleTaskDone()

	task.MarkTodo()

	assert.False(t, task.IsDone)
}

func newExampleTaskTodo() *task.Task {
	return &task.Task{
		ID: uuid.NewString(),
		Title: "basic exxample task",
		IsDone: false,
	}
}

func newExampleTaskDone() *task.Task {
	return &task.Task{
		ID: uuid.NewString(),
		Title: "basic exxample task",
		IsDone: true,
	}
}