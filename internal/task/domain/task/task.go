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

func (t *Task) Validate() error {
	if t.ID == "" {
		return errors.New("empty ID")
	}

	if t.Title == "" {
		return errors.New("empty Title")
	}
	
	return nil
}


type TaskService struct {
	list []Task
}

type PostTask struct {
	Title string `json:"title"`
}

func NewTask(taskData PostTask) (*Task, error){
	task := Task{
		ID: uuid.NewString(),
		Title: taskData.Title,
		IsDone: false,
	}

	err := task.Validate();
	if err != nil {
		return nil, err
	}
	return &task, nil
}


func MustNewTask(taskData PostTask) (*Task){

	newTask, err := NewTask(taskData)

	if err != nil {
		panic(err)
	}

	return newTask
}
