package service

import (
	"errors"
	"mime/multipart"
	"strconv"

	"github.com/ervinismu/devstore/internal/app/model"
	"github.com/ervinismu/devstore/internal/app/schema"
	"github.com/ervinismu/devstore/internal/pkg/reason"
	log "github.com/sirupsen/logrus"
)

type ImageUploader interface {
	UploadImage(input *multipart.FileHeader) (string, error)
}

type ProductService struct {
	productRepo  ProductRepository
	categoryRepo CategoryRepository
	uploader     ImageUploader
}

func NewProductService(productRepo ProductRepository, categoryRepo CategoryRepository, uploader ImageUploader) *ProductService {
	return &ProductService{
		productRepo:  productRepo,
		categoryRepo: categoryRepo,
		uploader:     uploader,
	}
}

// create product
func (cs *ProductService) Create(req *schema.CreateProductReq) error {
	var insertData model.Product

	insertData.Name = req.Name
	insertData.Description = req.Description
	insertData.Currency = req.Currency
	insertData.TotalStock = req.TotalStock
	insertData.IsActive = req.IsActive
	insertData.CategoryID = req.CategoryID
	insertData.Price = req.Price

	categoryID := strconv.Itoa(req.CategoryID)
	_, err := cs.categoryRepo.GetByID(categoryID)
	if err != nil {
		return errors.New(reason.CategoryNotFound)
	}

	// upload file to cloudinary
	imageURL, err := cs.uploader.UploadImage(req.Image)
	if err != nil {
		log.Error("upload image product : %w", err)
		return errors.New(reason.ProductCannotCreate)
	}

	insertData.ImageUrl = &imageURL

	// Return productID when create product
	err = cs.productRepo.Create(insertData)
	if err != nil {
		return errors.New(reason.ProductCannotCreate)
	}

	return nil
}

// get list product
func (cs *ProductService) BrowseAll(req *schema.BrowseProductReq) ([]schema.BrowseProductResp, error) {
	var resp []schema.BrowseProductResp

	dbSearch := model.BrowseProduct{}
	dbSearch.Page = req.Page
	dbSearch.PageSize = req.PageSize

	products, err := cs.productRepo.Browse(dbSearch)
	if err != nil {
		return nil, errors.New(reason.ProductCannotBrowse)
	}

	for _, value := range products {
		respData := schema.BrowseProductResp{
			ID:          value.ID,
			Name:        value.Name,
			Description: value.Description,
			Currency:    value.Currency,
			TotalStock:  value.TotalStock,
			IsActive:    value.IsActive,
			ImageURL:    value.ImageUrl,
			Price:       value.Price,
		}

		resp = append(resp, respData)
	}

	return resp, nil
}

// get detail product
func (cs *ProductService) GetByID(id string) (schema.DetailProductResp, error) {
	var resp schema.DetailProductResp

	product, err := cs.productRepo.GetByID(id)
	if err != nil {
		return resp, errors.New(reason.ProductCannotGetDetail)
	}

	categoryID := strconv.Itoa(product.CategoryID)
	category, err := cs.categoryRepo.GetByID(categoryID)
	if err != nil {
		return resp, errors.New(reason.ProductCannotGetDetail)
	}

	resp = schema.DetailProductResp{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Currency:    product.Currency,
		TotalStock:  product.TotalStock,
		IsActive:    product.IsActive,
		ImageURL:    product.ImageUrl,
		Price:       product.Price,
		Category: schema.Category{
			ID:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		},
	}

	return resp, nil
}

// update product by id
func (cs *ProductService) UpdateByID(id string, req *schema.UpdateProductReq) error {

	var updateData model.Product

	oldData, err := cs.productRepo.GetByID(id)
	if err != nil {
		return errors.New(reason.ProductNotFound)
	}

	updateData.ID = oldData.ID
	updateData.Name = req.Name
	updateData.Description = req.Description
	updateData.Currency = req.Currency
	updateData.TotalStock = req.TotalStock
	updateData.IsActive = req.IsActive
	updateData.CategoryID = req.CategoryID
	updateData.Price = req.Price

	err = cs.productRepo.Update(updateData)
	if err != nil {
		return errors.New(reason.ProductCannotUpdate)
	}

	return nil
}

// delete product by id
func (cs *ProductService) DeleteByID(id string) error {

	_, err := cs.productRepo.GetByID(id)
	if err != nil {
		return errors.New(reason.ProductNotFound)
	}

	err = cs.productRepo.DeleteByID(id)
	if err != nil {
		return errors.New(reason.ProductCannotDelete)
	}

	return nil
}
