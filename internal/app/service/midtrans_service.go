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

// Charge payment
func (svc *MidtransService) Charge(chargePayload schema.OrderChargePayload) error {
	// payload request
	payload, err := svc.payloadBCAVA(chargePayload)
	if err != nil {
		log.Error(fmt.Errorf("failed build charge payload: %w", err))
		return err
	}

	// charge request
	chargeResponse, err := svc.doMidtransChargeRequest("/v2/charge", payload)
	if err != nil {
		log.Error(fmt.Errorf("failed charge payment : %w", err))
		return err
	}

	// body response
	log.Info(fmt.Sprintf("change response TransactionStatus : %s", chargeResponse.TransactionStatus))
	log.Info(fmt.Sprintf("change response TransactionTime : %s", chargeResponse.TransactionTime))

	return nil
}

// process charge to midtrans
func (svc *MidtransService) doMidtransChargeRequest(path string, byteBody []byte) (schema.MidtransChargeResponse, error) {
	var (
		fullPath     = fmt.Sprintf("%s%s", svc.BaseURL, path)
		response     = schema.MidtransChargeResponse{}
		client       = &http.Client{Timeout: time.Duration(3) * time.Second}
		maxHTTPRetry = 3
		retryCount   = 1
		res          *http.Response
	)

	// init request
	req, err := http.NewRequest(http.MethodPost, fullPath, bytes.NewBuffer(byteBody))
	if err != nil {
		return response, err
	}

	// set header
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", svc.generateAuthHeader())

	// do request using http retry
	for {
		// process request
		res, err = client.Do(req)

		if err != nil || res.StatusCode != 200 {

			// check max retry
			if retryCount > maxHTTPRetry {
				return response, fmt.Errorf("HTTP max retry exceeded (%s) %d", req.URL, retryCount)
			}

			log.Error(fmt.Sprintf("HTTP error %d, retry request %d", res.StatusCode, retryCount))
			time.Sleep(time.Second * 1) // time 1 seconds every retry
			retryCount++                // increment retry count
			continue
		}
		break
	}

	defer res.Body.Close()

	// parse response body
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return response, fmt.Errorf("HTTP read response body %w", err)
	}

	json.Unmarshal(resBody, &response)

	return response, nil
}

// generate auth header
func (svc *MidtransService) generateAuthHeader() string {
	token := base64.StdEncoding.EncodeToString([]byte(svc.ServerKey))
	authHeader := fmt.Sprintf("Basic %s", token)
	return authHeader
}

// generate payload
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
		return nil, fmt.Errorf("payload bca va %w", err)
	}

	return byteBody, nil
}
