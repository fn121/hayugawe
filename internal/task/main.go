package main

import (
	"log"
	"net/http"
	"os"

	"github.com/fn121/hayugawe/internal/task/common/server"
	"github.com/fn121/hayugawe/internal/task/ports"
	"github.com/fn121/hayugawe/internal/task/service"
)

func RunHTTPServer(handler http.Handler) {
	log.Print("starting server...")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting port to %s", port)
	}

	log.Printf("listening on port %s", port)
	log.Panic(http.ListenAndServe(":"+port, handler))
}

func main() {
	app := service.NewApplication()
	apiMux := http.NewServeMux()
	handler := ports.HandlerFromMux(ports.NewHttpServer(app), apiMux)

	server.RunHTTPServer(handler)
}