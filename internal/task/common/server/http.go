package server

import (
	"log"
	"net/http"
	"os"
)

func RunHTTPServer(handler http.Handler) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting port to %s", port)
	}

	RunHTTPServerOnAddr(":"+port, handler)
}

func RunHTTPServerOnAddr(addr string, handler http.Handler) {
	log.Print("starting server...")

	// set base url
	mainMux := http.NewServeMux()
	mainMux.Handle("/api/", http.StripPrefix("/api", handler))

	log.Printf("listening on addr %s", addr)
	log.Panic(http.ListenAndServe(addr, mainMux))
}