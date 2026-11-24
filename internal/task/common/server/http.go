package server

import (
	"log"
	"net/http"
	"os"
)

func RunHTTPServer(handler http.Handler) {
	log.Print("starting server...")

	// set base url
	mainMux := http.NewServeMux()
	mainMux.Handle("/api/", http.StripPrefix("/api", handler))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting port to %s", port)
	}

	log.Printf("listening on port %s", port)
	log.Panic(http.ListenAndServe(":"+port, mainMux))
}