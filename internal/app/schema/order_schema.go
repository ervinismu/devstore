package schema

type OrderChargePayload struct {
	Amount        int
	OrderID       string
	CustomerName  string
	CustomerEmail string
	VANumber      string
}
