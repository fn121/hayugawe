package tests

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/fn121/hayugawe/internal/task/common/client/task"
	taskclient "github.com/fn121/hayugawe/internal/task/common/client/task"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type serverConfig struct {
	ServerUrl string
	hostport string
}

func NewServerConfig() *serverConfig {
	taskServiceHttpAddr := os.Getenv("TASK_SERVICE_HTTP_ADDR")
	if taskServiceHttpAddr == "" {
		panic("require env TASK_SERVICE_HTTP_ADDR")
	}

	return &serverConfig{
		ServerUrl: fmt.Sprintf("http://%s/api/", taskServiceHttpAddr),
		hostport: taskServiceHttpAddr,
	}
}

var (
	config = NewServerConfig()
)

type TaskHTTPClient struct {
	client *taskclient.ClientWithResponses
}

func NewTaskHTTPClient(t *testing.T) TaskHTTPClient {
	// @todo @fix Add token if needed with add optional args
	// see https://github.com/ThreeDotsLabs/wild-workouts-go-ddd-example/blob/8ecfcdf05b1462c4757bd2dcac9086c78e9f7791/internal/common/tests/clients.go#L39
	client, err := task.NewClientWithResponses(config.ServerUrl)
	require.NoError(t, err)

	return TaskHTTPClient{
		client: client,
	}
}

func (c TaskHTTPClient) CreateTask(t *testing.T, title string) (taskID string) {
	ctx := context.Background()
	resp, err := c.client.CreateTaskWithResponse(ctx, taskclient.CreateTaskJSONRequestBody{Title: title})
	require.NoError(t, err)

	contentLocation := resp.HTTPResponse.Header.Get("Content-Location")
	assert.NotEmpty(t, contentLocation)

	taskID = lastPathElement(contentLocation)
	assert.NotEmpty(t, taskID)

	return taskID
}

func (c TaskHTTPClient) GetTaskList(t *testing.T) []taskclient.Task  {
	ctx := context.Background()
	resp, err := c.client.GetTaskListWithResponse(ctx)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode())

	return *resp.JSON200
}

func (c TaskHTTPClient) MarkTaskAsDone(t *testing.T, taskID string) {
	ctx := context.Background()
	resp, err := c.client.MarkTaskAsDoneWithResponse(ctx, taskID)
	require.NoError(t, err)

	assert.Equal(t, http.StatusNoContent, resp.StatusCode())
}

func (c TaskHTTPClient) MarkTaskAsTodo(t *testing.T, taskID string) {
	ctx := context.Background()
	resp, err := c.client.MarkTaskAsTodoWithResponse(ctx, taskID)
	require.NoError(t, err)

	assert.Equal(t, http.StatusNoContent, resp.StatusCode())
}

func lastPathElement(path string) string {
	parts := strings.Split(path, "/")
	return parts[len(parts)-1]
}