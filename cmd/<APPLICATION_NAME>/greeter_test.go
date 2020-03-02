package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGreeter(t *testing.T) {
	srv := httptest.NewServer(Greeter("tester"))
	defer srv.Close()

	resp, err := http.Get(srv.URL)
	if resp != nil {
		defer func() { _ = resp.Body.Close() }()
	}

	if err != nil {
		t.Fatal("HTTP request to Greeter has failed:", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("HTTP server returned non-200 response, got %v", resp.StatusCode)
	}

	body := &Greeting{}

	if err := json.NewDecoder(resp.Body).Decode(body); err != nil {
		t.Fatal("Unable to parse response as JSON:", err)
	}

	if want, got := "tester", body.Name; want != got {
		t.Errorf("Name in the response does not match expected value, want: %#v, got: %#v", want, got)
	}

	if want, got := "Welcome, tester", body.Message; want != got {
		t.Errorf("Message in the response does not match expected value, want: %#v, got: %#v", want, got)
	}
}
