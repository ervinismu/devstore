package repository

import "github.com/ervinismu/devstore/internal/app/model"

type ICategoryRepository interface {
	Create(category model.Category) error
	Browse() ([]model.Category, error)
	Update(category model.Category) error
	GetByID(id string) (model.Category, error)
	DeleteByID(id string) error
}

type IProductRepository interface {
	Create(product model.Product) error
	Browse() ([]model.Product, error)
	GetByID(id string) (model.Product, error)
	Update(product model.Product) error
	DeleteByID(id string) error
}

type IUserRepository interface {
	Create(user model.User) error
	Browse() ([]model.User, error)
	GetByEmailAndUsername(email string, username string) (model.User, error)
	GetByEmail(email string) (model.User, error)
	Update(user model.User) error
	DeleteByID(id string) error
}

type IAuthRepository interface {
	Create(auth model.Auth) error
	DeleteAllByUserID(userID int) error
}
