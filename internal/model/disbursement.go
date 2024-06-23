package model

import "github.com/didikz/godisb/pkg/helpers"

const (
	BankAccountStatusVerified = 1
	BankAccountStatusBlocked  = 0

	DisbursementStatusPending = 0
	DisbursementStatusSuccess = 1
	DisbursementStatusFailed  = 2

	DisbursementStatusPendingString = "PENDING"
	DisbursementStatusSuccessString = "SUCCESS"
	DisbursementStatusFailedString  = "FAILED"
)

type User struct {
	ID       int64  `db:"id"`
	Email    string `db:"email"`
	Password string `db:"password"`
	Balance  int64  `db:"balance"`
}

type BankAccount struct {
	ID            int64  `db:"id"`
	Bank          string `db:"bank"`
	AccountNumber string `db:"account_number"`
	AccountName   string `db:"account_name"`
	Status        int    `db:"status"`
}

type Disbursement struct {
	ID              int64   `db:"id"`
	Bank            string  `db:"bank"`
	AccountNumber   string  `db:"account_number"`
	BeneficiaryName string  `db:"beneficiary_name"`
	Amount          int64   `db:"amount"`
	Remark          *string `db:"remark"`
	Status          int     `db:"status"`
	FailedNotes     *string `db:"failed_notes"`
	CreatedAt       int64   `db:"created_at"`
	FailedAt        *int64  `db:"failed_at"`
	CompletedAt     *int64  `db:"completed_at"`
	UserID          int64   `db:"user_id"`
	IdempotencyKey  string  `db:"idempotency_key"`
}

type CreateDisbursementPayload struct {
	Bank          string  `json:"bank"`
	AccountNumber string  `json:"account_number"`
	Amount        int64   `json:"amount"`
	Remark        *string `json:"remark"`
}

type DisbursementResponseObject struct {
	ID              int64  `json:"id"`
	Bank            string `json:"bank"`
	AccountNumber   string `json:"account_number"`
	BeneficiaryName string `json:"beneficiary_name"`
	Amount          int64  `json:"amount"`
	Remark          string `json:"remark"`
	Status          string `json:"status"`
	FailedNotes     string `json:"failed_notes"`
	CreatedAt       string `json:"created_at"`
	FailedAt        string `json:"failed_at"`
	CompletedAt     string `json:"completed_at"`
	IdempotencyKey  string `json:"idempotency_key"`
}

func (ba *BankAccount) IsBlocked() bool {
	return ba.Status == BankAccountStatusBlocked
}

func (d *Disbursement) parseStatusToString() string {
	switch d.Status {
	case DisbursementStatusSuccess:
		return DisbursementStatusSuccessString
	case DisbursementStatusFailed:
		return DisbursementStatusFailedString
	default:
		return DisbursementStatusPendingString
	}
}

func (d *Disbursement) parseCreatedAtToStringFormat() string {
	return helpers.UnixTimeToFormattedString(d.CreatedAt)
}

func (d *Disbursement) parseFailedAtToStringFormat() string {
	if d.FailedAt == nil {
		return ""
	}
	return helpers.UnixTimeToFormattedString(*d.FailedAt)
}

func (d *Disbursement) parseCompletedAtToStringFormat() string {
	if d.CompletedAt == nil {
		return ""
	}
	return helpers.UnixTimeToFormattedString(*d.CompletedAt)
}

func (d *Disbursement) ToJSONObject() DisbursementResponseObject {
	remark, failedNotes := "", ""
	if d.Remark != nil {
		remark = *d.Remark
	}
	if d.FailedNotes != nil {
		failedNotes = *d.FailedNotes
	}
	return DisbursementResponseObject{
		ID:              d.ID,
		Bank:            d.Bank,
		AccountNumber:   d.AccountNumber,
		BeneficiaryName: d.BeneficiaryName,
		Amount:          d.Amount,
		Remark:          remark,
		Status:          d.parseStatusToString(),
		FailedNotes:     failedNotes,
		CreatedAt:       d.parseCreatedAtToStringFormat(),
		FailedAt:        d.parseFailedAtToStringFormat(),
		CompletedAt:     d.parseCompletedAtToStringFormat(),
		IdempotencyKey:  d.IdempotencyKey,
	}
}
