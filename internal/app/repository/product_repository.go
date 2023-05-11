package repository

import (
	"errors"
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/ervinismu/devstore/internal/app/model"
	"github.com/jmoiron/sqlx"
)

type ProductRepository struct {
	DB *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) *ProductRepository {
	return &ProductRepository{DB: db}
}

// create product
func (cr *ProductRepository) Create(product model.Product) error {
	var (
		sqlStatement = `
			INSERT INTO products (name, description, currency, total_stock, is_active, category_id, image_url, price)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
			`
	)

	_, err := cr.DB.Exec(sqlStatement,
		product.Name,
		product.Description,
		product.Currency,
		product.TotalStock,
		product.IsActive,
		product.CategoryID,
		product.ImageUrl,
		product.Price,
	)
	if err != nil {
		log.Error(fmt.Errorf("error ProductRepository - Create : %w", err))
		return err
	}

	return nil
}

// get list product
func (cr *ProductRepository) Browse(search model.BrowseProduct) ([]model.Product, error) {
	var (
		limit        = search.PageSize
		offset       = limit * (search.Page - 1)
		products     []model.Product
		sqlStatement = `
			SELECT id, name, description, currency, total_stock, is_active, category_id, image_url, price
			FROM products
			LIMIT $1
			OFFSET $2
		`
	)

	rows, err := cr.DB.Queryx(sqlStatement, limit, offset)
	if err != nil {
		log.Error(fmt.Errorf("error ProductRepository - Browse : %w", err))
		return products, err
	}

	for rows.Next() {
		var product model.Product
		err := rows.StructScan(&product)
		if err != nil {
			log.Error(fmt.Errorf("error ProductRepository - Browse : %w", err))
		}
		products = append(products, product)
	}

	return products, nil
}

// get detail product
func (cr *ProductRepository) GetByID(id string) (model.Product, error) {
	var (
		sqlStatement = `
			SELECT id, name, description, currency, total_stock, is_active, category_id, image_url, price
			FROM products
			WHERE id = $1
		`
		product model.Product
	)
	err := cr.DB.QueryRowx(sqlStatement, id).StructScan(&product)
	if err != nil {
		log.Error(fmt.Errorf("error ProductRepository - GetByID : %w", err))
		return product, err
	}

	return product, nil
}

// update article by id
func (cr *ProductRepository) Update(product model.Product) error {
	var (
		sqlStatement = `
			UPDATE products
			SET updated_at = NOW(),
				name = $2,
				description = $3,
				currency = $4,
				total_stock = $5,
				is_active = $6,
				category_id	 = $7
				price	 = $8
			WHERE id = $1
		`
	)

	result, err := cr.DB.Exec(sqlStatement,
		product.ID,
		product.Name,
		product.Description,
		product.Currency,
		product.TotalStock,
		product.IsActive,
		product.CategoryID,
		product.Price,
	)
	if err != nil {
		log.Error(fmt.Errorf("error ProductRepository - Update : %w", err))
		return err
	}

	totalAffected, _ := result.RowsAffected()
	if totalAffected <= 0 {
		return errors.New("no record affected")
	}

	return nil
}

// delete article by id
func (cr *ProductRepository) DeleteByID(id string) error {
	var (
		sqlStatement = `
			DELETE FROM products
			WHERE id = $1
		`
	)

	result, err := cr.DB.Exec(sqlStatement, id)
	if err != nil {
		log.Error(fmt.Errorf("error ProductRepository - DeleteByID : %w", err))
		return err
	}

	totalAffected, _ := result.RowsAffected()
	if totalAffected <= 0 {
		return errors.New("no record affected")
	}

	return nil
}
