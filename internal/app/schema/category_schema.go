package schema

type GetCategoryResp struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CreateCategoryReq struct {
	Name        string
	Description string
}

type UpdateCategoryReq struct {
	Name        string
	Description string
}
