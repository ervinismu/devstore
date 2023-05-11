package model

type CartItem struct {
	ID         int `db:"id"`
	Quantity   int `db:"quantity"`
	BasePrice  int `db:"base_price"`
	TotalPrice int `db:"total_price"`
	CartID     int `db:"cart_id"`
	ProductID  int `db:"product_id"`
}
