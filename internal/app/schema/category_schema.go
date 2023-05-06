package schema

type BrowseCategoryReq struct {
	Page     int // Query number of pages
	PageSize int // search page size
}

type GetCategoryResp struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CreateCategoryReq struct {
	Name        string `validate:"required" json:"name"`
	Description string `validate:"required" json:"description"`
}

type UpdateCategoryReq struct {
	Name        string `validate:"required" json:"name"`
	Description string `validate:"required" json:"description"`
}
