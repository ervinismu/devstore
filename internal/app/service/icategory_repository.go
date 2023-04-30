package service

import "github.com/ervinismu/devstore/internal/app/model"

type CategoryRepository interface {
	Create(category model.Category) error
	Browse() ([]model.Category, error)
	Update(category model.Category) error
	GetByID(id string) (model.Category, error)
	DeleteByID(id string) error
}
