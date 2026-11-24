package service

import (
	"net/http"
	"os"
	"testing"

	"github.com/fn121/hayugawe/internal/task/ports"
)

func startService() (bool) {
	app := NewComponentTestApplication();
	mux := http.NewServeMux()
	taskServer := ports.NewHttpServer(app)
	taskServer.RegisterRoutes(mux)

	taskServiceHttpAddr := os.Getenv("TASK_SERVICE_HTTP_ADDR")
	if taskServiceHttpAddr == "" {
		panic("require env TASK_SERVICE_HTTP_ADDR")
	}
	
	go http.ListenAndServe(taskServiceHttpAddr, mux)
	return true
}

func TestMain(m *testing.M) {
	isRun := startService()
	if !isRun {
		os.Exit(1)
	}

	os.Exit(m.Run())
}