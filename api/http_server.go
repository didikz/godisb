package api

import (
	"encoding/json"
	"log"
	"net/http"

	httpapi "github.com/didikz/godisb/pkg/httpapi"
	m "github.com/didikz/godisb/pkg/httpapi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type HttpServer struct {
	listenAddress string
}

func NewHttpServer(listenAddress string) *HttpServer {
	return &HttpServer{
		listenAddress: listenAddress,
	}
}

func (s *HttpServer) Run() error {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)
	r.Use(m.HeaderValidator)

	r.Post("/disbursements", handleCreateDisbursement)

	log.Println("running server at", s.listenAddress)
	return http.ListenAndServe(s.listenAddress, r)
}

func handleCreateDisbursement(w http.ResponseWriter, r *http.Request) {
	dto := &CreateDisbursementPayload{}
	if err := json.NewDecoder(r.Body).Decode(dto); err != nil {
		httpapi.WriteJson(w, http.StatusInternalServerError, httpapi.GeneralResponseError{Error: err.Error()})
		return
	}

	remark := ""
	if dto.Remark != nil {
		remark = *dto.Remark
	}

	do := DisbursementResponseObject{
		ID:              100,
		Bank:            dto.Bank,
		AccountNumber:   dto.AccountNumber,
		BeneficiaryName: "John Doe",
		Amount:          dto.Amount,
		Remark:          remark,
		Status:          "SUCCESS",
		FailedNotes:     "",
		CreatedAt:       "2024-10-10 10:10:10",
		FailedAt:        "",
		CompletedAt:     "2024-10-10 10:10:10",
	}
	httpapi.WriteJson(w, http.StatusOK, do)
}
