package task

import (
	"github.com/google/uuid"
)

type Task struct {
	ID string `json:"id"`
	Title string `json:"title"`
	IsDone bool `json:"is_done"`
}

func(t *Task) MarkDone() {
	t.IsDone = true
}

func(t *Task) MarkTodo() {
	t.IsDone = false
}

func(t *Task) ToggleStatus() {
	t.IsDone = !t.IsDone
}


type TaskService struct {
	list []Task
}

type PostTask struct {
	Title string `json:"title"`
}

func NewTask(taskData PostTask) *Task{
	return &Task{
		ID: uuid.NewString(),
		Title: taskData.Title,
		IsDone: false,
	}	
}
