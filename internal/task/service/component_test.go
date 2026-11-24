package service

import (
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/fn121/hayugawe/internal/task/common/client/task"
	"github.com/fn121/hayugawe/internal/task/common/server"
	"github.com/fn121/hayugawe/internal/task/common/tests"
	"github.com/fn121/hayugawe/internal/task/ports"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateAndGetListTask(t *testing.T) {
	t.Parallel()
	
	client := tests.NewTaskHTTPClient(t)

	taskID := client.CreateTask(t, "do component test")
	require.NotEmpty(t, taskID)

	tasks := client.GetTaskList(t)

	taskIDs := []string{}

	for _, task := range tasks {
		taskIDs = append(taskIDs, task.Id)
	}

	assert.Contains(t, taskIDs, taskID)
}

func TestMarkTaskAsDone(t *testing.T) {
	t.Parallel()
	
	client := tests.NewTaskHTTPClient(t)

	taskID := client.CreateTask(t, "do component test")
	require.NotEmpty(t, taskID)

	client.MarkTaskAsDone(t, taskID)
	tasks := client.GetTaskList(t)

	var ourTask task.Task
	for _, task := range tasks {
		if task.Id == taskID {
			ourTask = task
		}
	}

	assert.Equal(t, true, ourTask.IsDone)
}

func startService() (bool) {
	taskServiceHttpAddr := os.Getenv("TASK_SERVICE_HTTP_ADDR")
	if taskServiceHttpAddr == "" {
		panic("require env TASK_SERVICE_HTTP_ADDR")
	}
	
	appTest := NewComponentTestApplication()
	mux := http.NewServeMux()
	handler := ports.HandlerFromMux(ports.NewHttpServer(appTest), mux)
	go server.RunHTTPServer(handler)

	ok := tests.WaitForPort(taskServiceHttpAddr)
	if !ok {
		log.Println("Timed out waiting for trainings HTTP to come up")
	}

	return ok
}

func TestMain(m *testing.M) {
	isRun := startService()
	if !isRun {
		os.Exit(1)
	}

	os.Exit(m.Run())
}