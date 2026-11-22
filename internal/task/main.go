package main

import (
	"log"
	"net/http"
	"os"

	"github.com/fn121/hayugawe/internal/task/app"
	"github.com/fn121/hayugawe/internal/task/ports"
)



func main() {
	log.Print("starting server...")
	mux := http.NewServeMux()

	service := app.TaskService{}
	httpServer := ports.NewHttpServer(service)
	httpServer.RegisterRoutes(mux)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting port to %s", port)
	}

	log.Printf("listening on port %s", port)
	http.ListenAndServe(":"+port, mux)
}