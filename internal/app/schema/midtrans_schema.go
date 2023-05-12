package schema

type MidtransChargeResponse struct {
	//midtrans
	Currency           string   `json:"currency,omitempty"`
	FraudStatus        string   `json:"fraud_status,omitempty"`
	GrossAmount        string   `json:"gross_amount,omitempty"`
	MerchantID         string   `json:"merchant_id,omitempty"`
	OrderID            string   `json:"order_id,omitempty"`
	PaymentType        string   `json:"payment_type,omitempty"`
	StatusCode         string   `json:"status_code,omitempty"`
	StatusMessage      string   `json:"status_message,omitempty"`
	TransactionID      string   `json:"transaction_id,omitempty"`
	TransactionStatus  string   `json:"transaction_status,omitempty"`
	BillerCode         string   `json:"biller_code,omitempty"`
	PaymentCode        string   `json:"payment_code,omitempty"`
	TransactionTime    string   `json:"transaction_time,omitempty"`
	ValidationMessages []string `json:"validation_messages,omitempty"`
	Actions            []struct {
		Method string `json:"method"`
		Name   string `json:"name"`
		URL    string `json:"url"`
	} `json:"actions,omitempty"`
}

