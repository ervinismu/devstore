package repository

import "github.com/ervinismu/devstore/internal/app/model"

type ICategoryRepository interface {
	BrowseAll() ([]model.Category, error)
	Create(category model.Category) error
	GetByID(id string) (model.Category, error)
}
