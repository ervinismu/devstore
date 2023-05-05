package schema

import "mime/multipart"

type BrowseProductReq struct {
	Page     int `json:"page" form:"page"`           // Query number of pages
	PageSize int `json:"page_size" form:"page_size"` // Search page size
}

type BrowseProductResp struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Currency    string  `json:"currency"`
	TotalStock  int     `json:"total_stock"`
	IsActive    bool    `json:"is_active"`
	ImageURL    *string `json:"image_url"`
}

type DetailProductResp struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Currency    string   `json:"currency"`
	TotalStock  int      `json:"total_stock"`
	IsActive    bool     `json:"is_active"`
	Category    Category `json:"category"`
	ImageURL    *string  `json:"image_url"`
}

type Category struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CreateProductReq struct {
	Name        string                `validate:"required" form:"name"`
	Description string                `validate:"required" form:"description"`
	Currency    string                `validate:"required" form:"currency"`
	TotalStock  int                   `validate:"required,number" form:"total_stock"`
	IsActive    bool                  `validate:"required,boolean" form:"is_active"`
	CategoryID  int                   `validate:"required,number" form:"category_id"`
	Image       *multipart.FileHeader `validate:"required,omitempty" form:"image"`
}

type UpdateProductReq struct {
	Name        string                `validate:"required" form:"name"`
	Description string                `validate:"required" form:"description"`
	Currency    string                `validate:"required" form:"currency"`
	TotalStock  int                   `validate:"required,number" form:"total_stock"`
	IsActive    bool                  `validate:"required,boolean" form:"is_active"`
	CategoryID  int                   `validate:"required,number" form:"category_id"`
	Image       *multipart.FileHeader `validate:"required,omitempty" form:"image"`
}
