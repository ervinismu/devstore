package service

import "github.com/ervinismu/devstore/internal/app/schema"

type ICategoryService interface {
	Create(req schema.CreateCategoryReq) error
	BrowseAll() ([]schema.GetCategoryResp, error)
	GetByID(id string) (schema.GetCategoryResp, error)
	UpdateByID(id string, req schema.UpdateCategoryReq) (error)
	DeleteByID(id string) error
}
