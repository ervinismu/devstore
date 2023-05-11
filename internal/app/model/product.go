package model

type Product struct {
	ID          int     `db:"id"`
	Name        string  `db:"name"`
	Description string  `db:"description"`
	Currency    string  `db:"currency"`
	Price       int     `db:"price"`
	TotalStock  int     `db:"total_stock"`
	IsActive    bool    `db:"is_active"`
	CategoryID  int     `db:"category_id"`
	ImageUrl    *string `db:"image_url"`
}

type BrowseProduct struct {
	Page     int
	PageSize int
}
