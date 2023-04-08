package repository

import "github.com/ervinismu/devstore/internal/app/model"

type ICategoryRepository interface {
	Create(category model.Category) error
	Browse() ([]model.Category, error)
	GetByID(id string) (model.Category, error)
	UpdateByID(id string, category model.Category) error
	DeleteByID(id string) error
}
