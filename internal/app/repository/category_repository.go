package repository

import (
	"fmt"
	"log"

	"github.com/ervinismu/devstore/internal/app/model"
	"github.com/jmoiron/sqlx"
)

type CategoryRepository struct {
	DB *sqlx.DB
}

func NewCategoryRepository(db *sqlx.DB) *CategoryRepository {
	return &CategoryRepository{DB: db}
}

// create category
func (cr *CategoryRepository) Create(category model.Category) error {
	var (
		sqlStatement = `
			INSERT INTO categories (name, description)
			VALUES ($1, $2)
			`
	)

	_, err := cr.DB.Exec(sqlStatement, category.Name, category.Description)
	if err != nil {
		log.Print(fmt.Errorf("error CategoryRepository - Create : %w", err))
		return err
	}

	return nil
}

// get list category
func (cr *CategoryRepository) Browse() ([]model.Category, error) {
	var (
		categories   []model.Category
		sqlStatement = `
			SELECT id, name, description
			FROM categories
		`
	)

	rows, err := cr.DB.Queryx(sqlStatement)
	if err != nil {
		log.Print(fmt.Errorf("error CategoryRepository - Browse : %w", err))
		return categories, err
	}

	for rows.Next() {
		var category model.Category
		rows.StructScan(&category)
		categories = append(categories, category)
	}

	return categories, nil
}

// get detail category
func (cr *CategoryRepository) GetByID(id string) (model.Category, error) {
	var (
		sqlStatement = `
			SELECT id, name, description
			FROM categories
			WHERE id = $1
		`
		category model.Category
	)
	err := cr.DB.QueryRowx(sqlStatement, id).StructScan(&category)
	if err != nil {
		log.Print(fmt.Errorf("error CategoryRepository - GetByID : %w", err))
		return category, err
	}

	return category, nil
}

// update article by id
// delete article by id
