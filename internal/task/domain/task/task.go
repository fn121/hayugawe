package task

import (
	"errors"

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

func NewTask(taskData PostTask) (*Task, error){
	if taskData.Title == "" {
		return nil, errors.New("empty Title")
	}

	return &Task{
		ID: uuid.NewString(),
		Title: taskData.Title,
		IsDone: false,
	}, nil
}


func MustNewTask(taskData PostTask) (*Task){

	newTask, err := NewTask(taskData)

	if err != nil {
		panic(err)
	}

	return newTask
}
