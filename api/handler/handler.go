package handler

import (
	"encoding/json"
	"net/http"

	"github.com/didikz/godisb/config"
	"github.com/didikz/godisb/internal/infrastructure"
	"github.com/didikz/godisb/internal/model"
	"github.com/didikz/godisb/internal/service"
	"github.com/didikz/godisb/internal/store"
	"github.com/didikz/godisb/pkg/httpapi"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
)

type handler struct {
	router              *chi.Mux
	disbursementService *service.DisbursementService
}

func NewHandler(router *chi.Mux, db *sqlx.DB, cfg config.Configuration) *handler {
	return &handler{
		router:              router,
		disbursementService: service.NewDisbursementService(*store.NewDisbursementRepository(db), *infrastructure.NewExternalApi(cfg)),
	}
}

func (h *handler) CreateDisbursement() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		dto := &model.CreateDisbursementPayload{}
		if err := json.NewDecoder(r.Body).Decode(dto); err != nil {
			httpapi.WriteJson(w, http.StatusInternalServerError, httpapi.GeneralResponseError{Error: err.Error()})
			return
		}

		validate := validator.New()
		err := validate.Struct(dto)
		if err != nil {
			httpapi.WriteJson(w, http.StatusUnprocessableEntity, httpapi.GeneralResponseError{Error: err.Error()})
			return
		}

		do, err := h.disbursementService.CreateDisbursement(ctx, *dto)
		if err != nil {
			httpapi.WriteJson(w, http.StatusInternalServerError, httpapi.GeneralResponseError{Error: err.Error()})
			return
		}

		httpapi.WriteJson(w, http.StatusOK, do)
	}
}
