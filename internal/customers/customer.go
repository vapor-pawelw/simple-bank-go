package customer

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	db "simple-bank/internal/db/generated"
)

func RegisterRoutes(r chi.Router, queries *db.Queries) {
	_ = queries
	r.Post("/customers", createCustomerHandler)
}

func createCustomerHandler(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusCreated)
}
