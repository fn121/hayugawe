package main

import (
	"log"
	"net/http"
	"os"

	"github.com/fn121/hayugawe/internal/task/ports"
	"github.com/fn121/hayugawe/internal/task/service"
)

func main() {
	log.Print("starting server...")
	mux := http.NewServeMux()

	app := service.NewApplication()
	
	httpServer := ports.NewHttpServer(app)
	httpServer.RegisterRoutes(mux)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting port to %s", port)
	}

	log.Printf("listening on port %s", port)
	log.Panic(http.ListenAndServe(":"+port, mux))
}