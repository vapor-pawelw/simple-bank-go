package app

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	customer "simple-bank/internal/customers"
	db "simple-bank/internal/db/generated"
	"simple-bank/internal/health"
)

func NewRouter(queries *db.Queries) http.Handler {
	r := chi.NewRouter()
	health.RegisterRoutes(r)
	customer.RegisterRoutes(r, queries)
	return r
}
