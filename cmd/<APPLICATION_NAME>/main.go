package main

import (
	"flag"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

var (
	commit  = "dev"
	version = "dev"
)

func main() {
	var addr string

	flag.StringVar(&addr, "addr", ":8080", "HTTP server bind address")
	flag.Parse()

	log.Printf("Starting <APPLICATION_NAME> version: %v, commit: %v", version, commit)

	router := mux.NewRouter()
	router.Use(Recover())

	router.Handle("/", Greeter("<APPLICATION_NAME>"))

	log.Printf("Starting server at %v...", addr)

	srv := &http.Server{
		Addr:              addr,
		Handler:           router,
		ReadTimeout:       5 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
		IdleTimeout:       5 * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("HTTP Server has failed: %v", err)
	}
}
