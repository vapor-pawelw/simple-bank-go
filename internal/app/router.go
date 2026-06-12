package app

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"simple-bank/internal/health"
)

func NewRouter() http.Handler {
	r := chi.NewRouter()
	health.RegisterRoutes(r)
	return r
}
