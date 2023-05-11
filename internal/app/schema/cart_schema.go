package schema

type AddToCartReq struct {
	ProductID int `validate:"required" json:"product_id"`
	Quantity  int    `validate:"required" json:"quantity"`
	UserID    int    `validate:"required" json:"user_id"`
}

type AddToCartResp struct {
}
