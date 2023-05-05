package service

import "github.com/ervinismu/devstore/internal/app/model"

type ProductRepository interface {
	Create(product model.Product) (int, error)
	Browse(search *model.BrowseProduct) ([]model.Product, error)
	GetByID(id string) (model.Product, error)
	UpdateImageUrl(id int, imageURL string) error
	Update(product model.Product) error
	DeleteByID(id string) error
}
