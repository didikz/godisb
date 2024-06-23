package service

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/didikz/godisb/internal/infrastructure"
	"github.com/didikz/godisb/internal/model"
	"github.com/didikz/godisb/internal/store"
	pkgmiddleware "github.com/didikz/godisb/pkg/httpapi/middleware"
)

type DisbursementService struct {
	DisbursementRepository store.DisbursementRepository
	ExternalApi            infrastructure.ExternalApi
}

func NewDisbursementService(disbursementRepository store.DisbursementRepository, externalApi infrastructure.ExternalApi) *DisbursementService {
	return &DisbursementService{
		DisbursementRepository: disbursementRepository,
		ExternalApi:            externalApi,
	}
}

func (s *DisbursementService) CreateDisbursement(ctx context.Context, p model.CreateDisbursementPayload) (*model.DisbursementResponseObject, error) {
	bankAccount, err := s.DisbursementRepository.GetBankAccount(ctx, p.Bank, p.AccountNumber)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("beneficiary bank account not found")
		}
		return nil, err
	}

	if bankAccount.IsBlocked() {
		return nil, fmt.Errorf("beneficiary bank account is blocked")
	}

	ik := ctx.Value(pkgmiddleware.CtxIdempotencyKey).(string)
	userID := int64(1)

	// if exists, return disbursement with existing idempotency key
	dis, err := s.DisbursementRepository.GetDisbursementByIdempotencyKey(ctx, ik, userID)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	// create & process disbursement
	if err != nil && err == sql.ErrNoRows {
		newdis, err := s.DisbursementRepository.CreateDisbursement(ctx, p, *bankAccount, ik, userID)
		if err != nil {
			return nil, err
		}
		// call bank partner
		status, failedCode := callBankPayment(ctx, &s.ExternalApi, newdis)
		var updateErr error
		if !status {
			updateErr = s.DisbursementRepository.UpdateFailedDisbursement(ctx, newdis.ID, model.DisbursementStatusFailed, failedCode)
		} else {
			updateErr = s.DisbursementRepository.UpdateSuccessDisbursement(ctx, newdis.ID, model.DisbursementStatusSuccess)
		}

		if updateErr != nil {
			fmt.Println(updateErr)
		}

		updatedDis, _ := s.DisbursementRepository.GetDisbursementById(ctx, newdis.ID)
		d := updatedDis.ToJSONObject()
		return &d, nil
	}

	d := dis.ToJSONObject()
	return &d, nil
}

func callBankPayment(ctx context.Context, ea *infrastructure.ExternalApi, dis *model.Disbursement) (bool, string) {
	switch dis.Bank {
	case "bca":
		dto := infrastructure.BCAPaymentPayload{
			TrxID:     dis.ID,
			AccNumber: dis.AccountNumber,
			Amount:    dis.Amount,
		}
		body, sc, err := ea.CallPaymentBCA(ctx, dto)
		if err != nil || sc != http.StatusOK {
			return false, ""
		}

		bcaResp := &infrastructure.BCASuccessResponse{}
		err = json.Unmarshal(body, bcaResp)
		if err != nil {
			return false, ""
		}
		if bcaResp.Code != "00" {
			return false, "0"
		}
		fmt.Println("bcaresp", bcaResp)
		return true, ""

	case "mandiri":
		dto := infrastructure.MandiriPaymentPayload{
			TrxID:     dis.ID,
			AccNumber: dis.AccountNumber,
			Amount:    dis.Amount,
		}
		body, sc, err := ea.CallPaymentMandiri(ctx, dto)
		if err != nil || sc != http.StatusOK {
			return false, "0"
		}

		mandiriResp := &infrastructure.MandiriFailedResponse{}
		err = json.Unmarshal(body, mandiriResp)
		if err != nil {
			return false, "0"
		}
		if mandiriResp.Code != "00" {
			return false, strconv.Itoa(mandiriResp.ReasonID)
		}
		return true, ""
	default:
		return false, "0"
	}
}
