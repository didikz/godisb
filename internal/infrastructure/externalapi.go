package infrastructure

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/didikz/godisb/config"
)

type BCAPaymentPayload struct {
	TrxID     int64  `json:"trx_id"`
	AccNumber string `json:"acccount_number"`
	Amount    int64  `json:"amount"`
}

type MandiriPaymentPayload struct {
	TrxID     int64  `json:"trx_id"`
	AccNumber string `json:"acccount_number"`
	Amount    int64  `json:"amount"`
}

type BCASuccessResponse struct {
	Code        string `json:"code"`
	ReferenceID string `json:"reference_id"`
}

type MandiriFailedResponse struct {
	Code        string `json:"code"`
	ReferenceID string `json:"reference_id"`
	ReasonID    int    `json:"reason_id"`
}

type ExternalApi struct {
	BCAConfig     config.BankAPIConfiguration
	MandiriConfig config.BankAPIConfiguration
}

func NewExternalApi(config config.Configuration) *ExternalApi {
	return &ExternalApi{
		BCAConfig:     config.ExternalApi.Bca,
		MandiriConfig: config.ExternalApi.Mandiri,
	}
}

func (a *ExternalApi) CallPaymentBCA(ctx context.Context, dto BCAPaymentPayload) ([]byte, int, error) {
	paymentUrl := fmt.Sprintf("%s%s", a.BCAConfig.BaseURL, "/bca/payment")
	buff, _ := json.Marshal(dto)
	req, _ := http.NewRequest(http.MethodPost, paymentUrl, bytes.NewBuffer(buff))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("X-Api-Key", a.BCAConfig.ApiKey)
	return invoke(req)
}

func (a *ExternalApi) CallPaymentMandiri(ctx context.Context, dto MandiriPaymentPayload) ([]byte, int, error) {
	paymentUrl := fmt.Sprintf("%s%s", a.BCAConfig.BaseURL, "/bca/payment")
	buff, _ := json.Marshal(dto)
	req, _ := http.NewRequest(http.MethodPost, paymentUrl, bytes.NewBuffer(buff))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("X-Api-Key", a.BCAConfig.ApiKey)
	return invoke(req)
}

func (m *MandiriFailedResponse) ParseFailedCode() string {
	switch m.ReasonID {
	case 0:
		return "CANT_RECEIVE_FUNDS"
	default:
		return "BANK_PARTNER_ERROR"
	}
}

func invoke(req *http.Request) ([]byte, int, error) {
	client := &http.Client{}
	response, error := client.Do(req)
	if error != nil {
		return nil, 500, error
	}
	defer response.Body.Close()

	fmt.Println("response Status:", response.Status)
	body, _ := io.ReadAll(response.Body)
	fmt.Println("response Body:", string(body))
	return body, response.StatusCode, nil
}
