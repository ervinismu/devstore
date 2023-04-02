package service

import "github.com/ervinismu/devstore/internal/app/schema"

type ICategoryService interface {
	BrowseAll() ([]schema.GetCategoryResp, error)
	Create(req schema.CreateCategoryReq) error
	GetByID(id string) (schema.GetCategoryResp, error)
}
