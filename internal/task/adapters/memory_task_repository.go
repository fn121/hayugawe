package adapters

import (
	"context"
	"errors"
	"fmt"

	"github.com/fn121/hayugawe/internal/task/domain/task"
)

type MemoryTaskRepository struct {
	list []task.Task
}

// Find implements [task.Repository].
func (mtr *MemoryTaskRepository) Find(_ context.Context ) ([]task.Task, error) {
	return mtr.list, nil
}

// FindByID implements [task.Repository].
func (mtr *MemoryTaskRepository) FindByID(_ context.Context, taskID string) (*task.Task, error) {
	for idx, t := range mtr.list {
		if t.ID == taskID {
			return &mtr.list[idx], nil
		}
	}
	return nil, fmt.Errorf("task with id %q not found", taskID)
}

// Save implements [task.Repository].
func (mtr *MemoryTaskRepository) Save(_ context.Context, task task.Task) error {
	ctx := context.Background()
	existTask, err := mtr.FindByID(ctx, task.ID)
	
	// task not found, add new task
	if err != nil {
		mtr.list = append(mtr.list, task)
	} else {
		// update task
		existTask.Title = task.Title
		existTask.IsDone = task.IsDone
	}
	return nil
}

func (mtr *MemoryTaskRepository) Insert(_ context.Context, task task.Task) error {
	ctx := context.Background()
	existTask, _ := mtr.FindByID(ctx, task.ID)

	if existTask != nil {
		return errors.New("task already exist")
	} else {
		mtr.list = append(mtr.list, task)
		return nil
	}
}

// UpdateByID implements [task.Repository].
func (mtr *MemoryTaskRepository) UpdateByID(_ context.Context, taskID string, task task.Task) error {
	ctx := context.Background()
	existTask, err := mtr.FindByID(ctx, task.ID)
	
	// task not found, add new task
	if err != nil {
		return fmt.Errorf("fail to find task: %w", err)
	}

	// update task
	existTask.Title = task.Title
	existTask.IsDone = task.IsDone
	return nil
}

func NewMemoryTaskRepository() MemoryTaskRepository {
	newTask :=task.Task{ID: "cbb87a8c-c30d-4f01-be30-83181b3d9b41", Title: "Makan siang", IsDone: false}
	sampleTasks := []task.Task{newTask}
	
	return MemoryTaskRepository{
		list: sampleTasks,
	}
}