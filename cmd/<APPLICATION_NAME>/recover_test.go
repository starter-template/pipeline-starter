package main

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecover(t *testing.T) {
	middleware := Recover()

	srv := httptest.NewServer(middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		panic(errors.New("something went very wrong"))
	})))

	defer srv.Close()

	resp, err := http.Get(srv.URL)
	if resp != nil {
		defer func() { _ = resp.Body.Close() }()
	}

	if err != nil {
		t.Fatal("HTTP request to the server has failed:", err)
	}

	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("HTTP server returned non-500 response, got %v", resp.StatusCode)
	}
}

