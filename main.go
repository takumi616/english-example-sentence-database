package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	err := http.ListenAndServe(
		":8000",
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Request Path: %s", r.URL.Path[1:])
		}),
	)
	if err != nil {
		log.Fatalf("Failed to run http server: %v", err)
	}
}