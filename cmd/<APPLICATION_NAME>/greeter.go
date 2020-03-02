package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Greeting struct {
	Name    string
	Message string
}

func Greeter(name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		greeting := Greeting{
			Name:    name,
			Message: fmt.Sprintf("Welcome, %v", name),
		}

		if err := json.NewEncoder(w).Encode(greeting); err != nil {
			panic(err)
		}
	})
}
