package schema

type OrderReq struct {
	UserID int
}

type OrderChargePayload struct {
	Amount        int
	OrderID       string
	CustomerName  string
	CustomerEmail string
	VANumber      string
}
