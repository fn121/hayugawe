package app

import (
	"context"
	"fmt"

	"github.com/fn121/hayugawe/internal/task/domain/task"
)

type TaskService struct {
	repo task.Repository
}

func NewTaskService(repo task.Repository) TaskService {
	return TaskService{repo}
}

func (ts *TaskService) GetTask(taskID string) (*task.Task, error) {
	ctx := context.Background()
	return ts.repo.FindByID(ctx, taskID)
}

func (ts *TaskService) GetTaskList() ([]task.Task, error) {
	ctx := context.Background()
	tasks, err := ts.repo.Find(ctx)

	return tasks, err
}

func (ts *TaskService) CreateTask(taskData task.PostTask) (error) {
	ctx := context.Background()
	
	newTask, err := task.NewTask(taskData)
	if err != nil {
		return fmt.Errorf("fail to create task: %w", err)
	}

	errSave := ts.repo.Save(ctx, *newTask)
	if errSave != nil {
		return fmt.Errorf("fail to save: %w", errSave)
	}
	return nil
}

func (ts *TaskService) MarkTaskAsTodo(taskID string) (error) {
	ctx := context.Background()

	task, err := ts.GetTask(taskID)
	if err != nil {
		return fmt.Errorf("fail to get task: %w", err)
	}

	task.MarkTodo()
	ts.repo.UpdateByID(ctx, taskID, *task)

	return nil
}

func (ts *TaskService) MarkTaskAsDone(taskID string) (error) {
	ctx := context.Background()

	task, err := ts.GetTask(taskID)
	if err != nil {
		return fmt.Errorf("fail to get task: %w", err)
	}

	task.MarkDone()
	ts.repo.UpdateByID(ctx, taskID, *task)

	return nil
}