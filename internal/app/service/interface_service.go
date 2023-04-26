package service

import "github.com/ervinismu/devstore/internal/app/schema"

type ICategoryService interface {
	Create(req *schema.CreateCategoryReq) error
	BrowseAll() ([]schema.GetCategoryResp, error)
	GetByID(id string) (schema.GetCategoryResp, error)
	UpdateByID(id string, req *schema.UpdateCategoryReq) error
	DeleteByID(id string) error
}

type IProductService interface {
	Create(req *schema.CreateProductReq) error
	BrowseAll() ([]schema.BrowseProductResp, error)
	GetByID(id string) (schema.DetailProductResp, error)
	UpdateByID(id string, req *schema.UpdateProductReq) error
	DeleteByID(id string) error
}

type IRegistrationService interface {
	Register(req *schema.RegisterReq) error
}
