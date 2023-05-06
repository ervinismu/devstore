package schema

import "mime/multipart"

type BrowseProductReq struct {
	Page     int // Query number of pages
	PageSize int // search page size
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
	ImageURL    *string  `json:"image_url"`
	Category    Category `json:"category"`
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
	Image       *multipart.FileHeader `form:"image" validate:"required"`
}

type UpdateProductReq struct {
	Name        string                `validate:"required" form:"name"`
	Description string                `validate:"required" form:"description"`
	Currency    string                `validate:"required" form:"currency"`
	TotalStock  int                   `validate:"required,number" form:"total_stock"`
	IsActive    bool                  `validate:"required,boolean" form:"is_active"`
	CategoryID  int                   `validate:"required,number" form:"category_id"`
	Image       *multipart.FileHeader `form:"image" validate:"required"`
}
