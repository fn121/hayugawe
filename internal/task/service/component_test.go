package service

import (
	"net/http"
	"os"
	"testing"
)

func startService() (bool) {
	mux := http.NewServeMux()

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