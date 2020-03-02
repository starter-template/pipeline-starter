package main

import (
	"log"
	"net/http"
)

func Recover() func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					log.Printf("HTTP handler triggered a panic: %v", err)

					w.WriteHeader(http.StatusInternalServerError)
				}
			}()

			h.ServeHTTP(w, r)
		})
	}
}
