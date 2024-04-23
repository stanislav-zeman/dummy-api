package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type Status struct {
	Status string
	Time   time.Time
}

func main() {
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
	http.ListenAndServe(":8080", nil)
}
