package api

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func API() *chi.Mux {
	a := chi.NewRouter()

	a.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	return a
}
