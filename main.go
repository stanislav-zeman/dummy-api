package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"
)

type Message struct {
	Message string
}

type Status struct {
	Status string
	Time   time.Time
}

func main() {
	http.HandleFunc(
		"/",
		func(w http.ResponseWriter, r *http.Request) {
			m := Message{
				Message: "Hello world!",
			}

			bytes, err := json.Marshal(m)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			w.Write(bytes)
		},
	)
	http.HandleFunc(
		"/status",
		func(w http.ResponseWriter, r *http.Request) {
			s := Status{
				Status: "OK",
				Time:   time.Now(),
			}

			bytes, err := json.Marshal(s)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			w.Write(bytes)
		},
	)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	hostname := fmt.Sprintf(":%s", port)
	slog.Info("starting server...", "hostname", hostname)
	http.ListenAndServe(hostname, nil)
}
