package api

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
}
