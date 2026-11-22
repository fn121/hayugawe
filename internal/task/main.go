package main

import (
	"log"
	"net/http"
	"os"

	"github.com/fn121/hayugawe/internal/task/adapters"
	"github.com/fn121/hayugawe/internal/task/app"
	"github.com/fn121/hayugawe/internal/task/ports"
)



func main() {
	log.Print("starting server...")
	mux := http.NewServeMux()

	repository := adapters.NewMemoryTaskRepository()
	service := app.NewTaskService(&repository)
	httpServer := ports.NewHttpServer(service)
	httpServer.RegisterRoutes(mux)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting port to %s", port)
	}

	log.Printf("listening on port %s", port)
	log.Panic(http.ListenAndServe(":"+port, mux))
}