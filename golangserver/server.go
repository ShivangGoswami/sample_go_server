package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/rs/cors"
)

type Response struct {
	TimeStamp time.Time `json:"timestamp"`
	Greetings string    `json:"greeting"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Response{time.Now(), "Hi there!"})
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("../testApp/dist/testApp")))
	mux.HandleFunc("/endpoint", handler)
	handler := cors.Default().Handler(mux)
	http.ListenAndServe(":1234", handler)
	// http.HandleFunc("/endpoint", handler)
	// log.Fatal(http.ListenAndServe(":1234", nil))
}
