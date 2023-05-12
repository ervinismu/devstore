package service

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/ervinismu/devstore/internal/app/schema"
	log "github.com/sirupsen/logrus"
)

type MidtransService struct {
	ServerKey  string
	MerchantID string
	BaseURL    string
}

func NewMidtransService(serverKey string, merchantID string, baseURL string) *MidtransService {
	return &MidtransService{
		ServerKey:  serverKey,
		MerchantID: merchantID,
		BaseURL:    baseURL,
	}
}

func (svc *MidtransService) Charge(paymentPayload schema.OrderChargePayload) error {
	payload, err := svc.payloadBCAVA(paymentPayload)
	if err != nil {
		log.Error(fmt.Errorf("failed build payload charge payment : %w", err))
		return err
	}

	chargeResponse, err := svc.doMidtransChargeRequest("/v2/charge", payload)
	if err != nil {
		log.Error(fmt.Errorf("failed charge payment : %w", err))
		return err
	}

	log.Info(chargeResponse)
	log.Info(fmt.Sprintf("chargeResponse.TransactionStatus : %s", chargeResponse.TransactionStatus))
	log.Info(fmt.Sprintf("chargeResponse.TransactionTime : %s", chargeResponse.TransactionTime))
	log.Info(fmt.Sprintf("chargeResponse.OrderID : %s", chargeResponse.OrderID))

	return nil
}

func (svc *MidtransService) doMidtransChargeRequest(path string, byteBody []byte) (schema.MidtransChargeResponse, error) {
	var (
		fullPath     = fmt.Sprintf("%s%s", svc.BaseURL, path)
		response     = schema.MidtransChargeResponse{}
		client       = &http.Client{Timeout: time.Duration(3) * time.Second}
		maxHTTPRetry = 3
		retryCount   = 1
		res          *http.Response
	)

	req, err := http.NewRequest(http.MethodPost, fullPath, bytes.NewBuffer(byteBody))
	if err != nil {
		return response, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", svc.GenerateAuthHeader())

	// request with retry
	for {
		res, err = client.Do(req)
		if err != nil || res.StatusCode != 200 {

			// check max retry exceeded
			if retryCount > maxHTTPRetry {
				return response, fmt.Errorf("HTTP max retry exceeded (%s) %d", req.URL, retryCount)
			}

			log.Error(fmt.Sprintf("HTTP error %d, retry request (%s) %dx", res.StatusCode, req.URL, retryCount))
			time.Sleep(time.Second * 1) // sleep 1 seconds
			retryCount++                // increment retry count
			continue
		}
		break
	}

	defer res.Body.Close()

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return response, fmt.Errorf("HTTP read response body (%w)", err)
	}

	json.Unmarshal(resBody, &response)

	return response, nil
}

func (svc *MidtransService) GenerateAuthHeader() string {
	token := base64.StdEncoding.EncodeToString([]byte(svc.ServerKey))
	authHeader := fmt.Sprintf("Basic %s", token)

	return authHeader
}

// charge BCA VA
func (svc *MidtransService) payloadBCAVA(data schema.OrderChargePayload) ([]byte, error) {
	payload := map[string]interface{}{
		"payment_type": "bank_transfer",
		"transaction_details": map[string]interface{}{
			"gross_amount": data.Amount,
			"order_id":     data.OrderID,
		},
		"customer_details": map[string]interface{}{
			"email":      data.CustomerEmail,
			"first_name": data.CustomerName,
		},
		"bank_transfer": map[string]interface{}{
			"bank":      "bca",
			"va_number": data.VANumber,
		},
	}

	byteBody, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("payload bca va : %w", err)
	}

	return byteBody, nil
}
