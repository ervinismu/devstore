package schema

type CreateCategoryReq struct {
	Name        string
	Description string
}

type GetCategoryResp struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

