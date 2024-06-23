package store

import (
	"context"
	"time"

	"github.com/didikz/godisb/internal/model"
	"github.com/jmoiron/sqlx"
)

const (
	queryGetBankAccount      = `SELECT id, bank, account_number, account_name, status FROM bank_accounts WHERE bank =? AND account_number =?`
	queryGetDisbursementById = `SELECT id, bank, account_number, beneficiary_name, amount, remark, status, failed_notes, created_at, failed_at, completed_at, idempotency_key 
											FROM disbursements WHERE id=?`
	queryGetDisbursementByIdempotencyKey = `SELECT id, bank, account_number, beneficiary_name, amount, remark, status, failed_notes, created_at, failed_at, completed_at, idempotency_key 
											FROM disbursements WHERE idempotency_key =? AND user_id =?`
	queryInsertPendingDisbursement = `INSERT INTO disbursements (bank, account_number, beneficiary_name, amount, remark, status, created_at, idempotency_key, user_id) 
									VALUES (:bank, :account_number, :beneficiary_name, :amount, :remark, :status, :created_at, :idempotency_key, :user_id)`
	queryUpdateSuccessDisbursement = `UPDATE disbursements SET status =?, completed_at=? WHERE id = ?`
	queryUpdateFailedDisbursement  = `UPDATE disbursements SET status =?, failed_at=?, failed_notes=? WHERE id = ?`
)

type DisbursementRepository struct {
	DB *sqlx.DB
}

func NewDisbursementRepository(db *sqlx.DB) *DisbursementRepository {
	return &DisbursementRepository{
		DB: db,
	}
}

func (r *DisbursementRepository) CreateDisbursement(ctx context.Context, p model.CreateDisbursementPayload, ba model.BankAccount, ik string, userID int64) (*model.Disbursement, error) {
	dis := model.Disbursement{
		Bank:            p.Bank,
		AccountNumber:   p.AccountNumber,
		BeneficiaryName: ba.AccountName,
		Amount:          p.Amount,
		Remark:          p.Remark,
		Status:          model.DisbursementStatusPending,
		CreatedAt:       time.Now().Unix(),
		IdempotencyKey:  ik,
		UserID:          userID,
	}
	res, err := r.DB.NamedExec(queryInsertPendingDisbursement, &dis)
	if err != nil {
		return nil, err
	}
	id, _ := res.LastInsertId()
	dis.ID = id
	return &dis, nil
}

func (r *DisbursementRepository) UpdateSuccessDisbursement(ctx context.Context, id int64, status int) error {
	updatedAt := time.Now().Unix()
	_, err := r.DB.ExecContext(ctx, queryUpdateSuccessDisbursement, status, updatedAt, id)
	return err
}

func (r *DisbursementRepository) UpdateFailedDisbursement(ctx context.Context, id int64, status int, failedNotes string) error {
	failedAt := time.Now().Unix()
	_, err := r.DB.ExecContext(ctx, queryUpdateFailedDisbursement, status, failedAt, failedNotes, id)
	return err
}

func (r *DisbursementRepository) GetDisbursementById(ctx context.Context, id int64) (*model.Disbursement, error) {
	d := model.Disbursement{}
	if err := r.DB.Get(&d, queryGetDisbursementById, id); err != nil {
		return nil, err
	}
	return &d, nil
}

func (r *DisbursementRepository) GetDisbursementByIdempotencyKey(ctx context.Context, ik string, userID int64) (*model.Disbursement, error) {
	d := model.Disbursement{}
	if err := r.DB.Get(&d, queryGetDisbursementByIdempotencyKey, ik, userID); err != nil {
		return nil, err
	}
	return &d, nil
}

func (r *DisbursementRepository) GetUser(id int64) (*model.User, error) {
	var u *model.User
	return u, nil
}

func (r *DisbursementRepository) GetBankAccount(ctx context.Context, bank string, accountNumber string) (*model.BankAccount, error) {
	b := model.BankAccount{}
	if err := r.DB.Get(&b, queryGetBankAccount, bank, accountNumber); err != nil {
		return nil, err
	}
	return &b, nil
}
