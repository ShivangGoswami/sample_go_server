package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type Response struct {
	TimeStamp time.Time `json:"timestamp"`
	Greetings string    `json:"greeting"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Response{time.Now(), "Hi there!"})
}

func main() {
	http.HandleFunc("/endpoint", handler)
	log.Fatal(http.ListenAndServe(":1234", nil))
}
