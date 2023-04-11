package schema

type BrowseProductResp struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Currency    string `json:"currency"`
	TotalStock  int    `json:"total_stock"`
	IsActive    bool   `json:"is_active"`
}

type DetailProductResp struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Currency    string   `json:"currency"`
	TotalStock  int      `json:"total_stock"`
	IsActive    bool     `json:"is_active"`
	Category    Category `json:"category"`
}

type Category struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CreateProductReq struct {
	Name        string `validate:"required" json:"name"`
	Description string `validate:"required" json:"description"`
	Currency    string `validate:"required" json:"currency"`
	TotalStock  int    `validate:"required,number" json:"total_stock"`
	IsActive    bool   `validate:"required,boolean" json:"is_active"`
	CategoryID  int    `validate:"required,number" json:"category_id"`
}

type UpdateProductReq struct {
	Name        string `validate:"required" json:"name"`
	Description string `validate:"required" json:"description"`
	Currency    string `validate:"required" json:"currency"`
	TotalStock  int    `validate:"required,number" json:"total_stock"`
	IsActive    bool   `validate:"required,boolean" json:"is_active"`
	CategoryID  int    `validate:"required,number" json:"category_id"`
}
