package main

import (
	"fmt"
	"net/http"
	"time"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func timeNow(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Agora são: %s", time.Now())
}

func main() {
	http.HandleFunc("/", greet)
	http.HandleFunc("/time", timeNow)
	http.ListenAndServe(":8080", nil)
}