package handler

import (
	"github.com/didikz/godisb/config"
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

func RegisterHandler(router *chi.Mux, db *sqlx.DB, cfg config.Configuration) {
	handler := NewHandler(router, db, cfg)

	router.Post("/api/v1/disbursements", handler.CreateDisbursement())
}
