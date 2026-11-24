package adapters

import (
	"context"
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

func domainToModel(t task.Task) (TaskModel) {
	return TaskModel{ID: t.ID, Title: t.Title, IsDone: t.IsDone}
}

// Find implements [task.Repository].
func (mtr *MemoryTaskRepository) FindTask(_ context.Context ) ([]task.Task, error) {
	tasks := mtr.list.ToDomain()
	return tasks, nil
}

// FindTaskByID implements [task.Repository].
func (mtr *MemoryTaskRepository) FindTaskByID(_ context.Context, taskID string) (*task.Task, error) {
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
func (mtr *MemoryTaskRepository) AddTask(_ context.Context, task task.Task) error {
	ctx := context.Background()
	existTask, err := mtr.FindTaskByID(ctx, task.ID)
	
	// task not found, add new task
	if err != nil {
		mtr.list = append(mtr.list, domainToModel(task))
	} else {
		// update task
		existTask.Title = task.Title
		existTask.IsDone = task.IsDone
	}
	return nil
}

// UpdateByID implements [task.Repository].
func (mtr *MemoryTaskRepository) UpdateTask(_ context.Context, taskID string, t task.Task) error {
	err := t.Validate()
	if err != nil {
		return err
	}
	t.ID = taskID

	// update
	for i := range mtr.list {
		if mtr.list[i].ID == taskID {
			mtr.list[i] = domainToModel(t)
			return nil
		}
	}
	return task.ErrTaskNotFound
}

func NewMemoryTaskRepository(tasklist TaskList) *MemoryTaskRepository {
	
	return &MemoryTaskRepository{
		list: tasklist,
	}
}

func NewMemoryTaskRepositoryWithSample() *MemoryTaskRepository {
	newTask :=TaskModel{ID: "cbb87a8c-c30d-4f01-be30-83181b3d9b41", Title: "Makan siang", IsDone: false}
	sampleTasks := []TaskModel{newTask}
	
	return NewMemoryTaskRepository(sampleTasks)
}

// UpdateByID implements [query.GetTaskListReadModel].
func (mtr MemoryTaskRepository) GetTaskList() ([]query.Task, error) {
	return mtr.list.ToReadModel(), nil
}