package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator"
	"github.com/takumi616/go-english-vocabulary-api/handler"
	"github.com/takumi616/go-english-vocabulary-api/store"
)

// Not to depend on internal implementation,
// return http.Handler interface instead of type http.ServeMux.
func SetUpRouting() http.Handler {
	mux := chi.NewRouter()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"Status": "ok"}`))
	})

	v := validator.New()
	av := &handler.AddVocabulary{
		Store:     store.Vocabularies,
		Validator: v,
	}
	mux.Post("/vocabularies", av.ServeHTTP)
	return mux
}
