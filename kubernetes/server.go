package main

import (
	"fmt"
	"net/http"
	"os"
)

var isHealthy = true

func main() {
	http.HandleFunc("/healthz", Healthz)
	http.HandleFunc("/set-unhealth", SetUnhealth)
	http.HandleFunc("/", Hello)
	http.ListenAndServe(":8000", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("NAME")
	fmt.Fprintf(w, "Hello World! %s", name)
}

func Healthz(w http.ResponseWriter, r *http.Request) {
	if isHealthy {
		w.WriteHeader(200)
		w.Write([]byte("ok"))
	} else {
		w.WriteHeader(500)
		w.Write([]byte("Is unhealth"))
	}
}

func SetUnhealth(w http.ResponseWriter, r *http.Request) {
	isHealthy = false

	w.WriteHeader(201)
	w.Write([]byte("Now is unhealth"))
}
