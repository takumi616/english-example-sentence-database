package main

import (
	"net/http"
)

// Not to depend on internal implementation,
// return http.Handler interface instead of type http.ServeMux.
func SetUpRouting() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"Status": "ok"}`))
	})

	return mux
}
