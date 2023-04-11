package model

type Product struct {
	ID          int    `db:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`
	Currency    string `db:"currency"`
	TotalStock  int    `db:"total_stock"`
	IsActive    bool   `db:"is_active"`
	CategoryID  int    `db:"category_id"`
}
