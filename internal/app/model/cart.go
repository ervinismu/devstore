package model

type Cart struct {
	ID         int `db:"id"`
	UserID     int `db:"user_id"`
	TotalPrice int `db:"total_price"`
}
