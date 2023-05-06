package service

import "github.com/ervinismu/devstore/internal/app/model"

type ProductRepository interface {
	Create(product model.Product) error
	Browse(search model.BrowseProduct) ([]model.Product, error)
	GetByID(id string) (model.Product, error)
	Update(product model.Product) error
	DeleteByID(id string) error
}
