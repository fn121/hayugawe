package adapters

import (
	"context"
	"errors"
	"fmt"

	"github.com/fn121/hayugawe/internal/task/app/query"
	"github.com/fn121/hayugawe/internal/task/domain/task"
)

type TaskModel struct {
	ID string `json:"id"`
	Title string `json:"title"`
	IsDone bool `json:"is_done"`
}

func (t TaskModel) TypeDomain() task.Task {
	return task.Task{ID: t.ID, Title: t.Title, IsDone: t.IsDone}
}

func (t TaskModel) ReadModel() query.Task {
	return query.Task{ID: t.ID, Title: t.Title, IsDone: t.IsDone}
}

type TaskList []TaskModel

func (tl TaskList) ToDomain() (tasks []task.Task) {
	for _, t := range tl {
		tasks = append(tasks, t.TypeDomain())
	}
	return tasks
}

func (tl TaskList) ToReadModel() (tasks []query.Task) {
	for _, t := range tl {
		tasks = append(tasks, t.ReadModel())
	}
	return tasks
}

type MemoryTaskRepository struct {
	list TaskList
}

func (mtr *MemoryTaskRepository) unmarshalListToDomain( ) ([]task.Task) {
	tasks := []task.Task{}

	for _, t := range mtr.list {
		tasks = append(tasks, t.TypeDomain())
	}
	
	return tasks
}

func (mtr *MemoryTaskRepository) marshalItemFromDomain(t task.Task) (TaskModel) {
	return TaskModel{ID: t.ID, Title: t.Title, IsDone: t.IsDone}
}

// Find implements [task.Repository].
func (mtr *MemoryTaskRepository) Find(_ context.Context ) ([]task.Task, error) {
	tasks := mtr.unmarshalListToDomain()
	return tasks, nil
}

// FindByID implements [task.Repository].
func (mtr *MemoryTaskRepository) FindByID(_ context.Context, taskID string) (*task.Task, error) {
	for idx, t := range mtr.list {
		if t.ID == taskID {
			task := mtr.list[idx]
			taskDomain := task.TypeDomain()
			return  &taskDomain , nil
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
		mtr.list = append(mtr.list, mtr.marshalItemFromDomain(task))
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
		mtr.list = append(mtr.list, mtr.marshalItemFromDomain(task))
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
	newTask :=TaskModel{ID: "cbb87a8c-c30d-4f01-be30-83181b3d9b41", Title: "Makan siang", IsDone: false}
	sampleTasks := []TaskModel{newTask}
	
	return MemoryTaskRepository{
		list: sampleTasks,
	}
}