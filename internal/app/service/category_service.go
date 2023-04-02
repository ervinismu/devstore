package service

import (
	"errors"

	"github.com/ervinismu/devstore/internal/app/model"
	"github.com/ervinismu/devstore/internal/app/repository"
	"github.com/ervinismu/devstore/internal/app/schema"
)

type CategoryService struct {
	repo repository.ICategoryRepository
}

func NewCategoryService(repo repository.ICategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

// get all categories
func (cs *CategoryService) BrowseAll() ([]schema.GetCategoryResp, error) {
	var resp []schema.GetCategoryResp

	categories, err := cs.repo.BrowseAll()
	if err != nil {
		return nil, errors.New("cannot get categories data")
	}

	for _, value := range categories {
		var respData schema.GetCategoryResp
		respData.Description = value.Description
		respData.Name = value.Name
		respData.ID = value.ID
		resp = append(resp, respData)
	}

	return resp, nil
}

// create category
func (cs *CategoryService) Create(req schema.CreateCategoryReq) error {
	var insertData model.Category

	insertData.Name = req.Name
	insertData.Description = req.Description

	err := cs.repo.Create(insertData)
	if err != nil {
		return errors.New("cannot create category")
	}

	return nil
}

// get one category by id
func (cs *CategoryService) GetByID(id string) (schema.GetCategoryResp, error) {
	var resp schema.GetCategoryResp

	category, err := cs.repo.GetByID(id)
	if err != nil {
		return resp, errors.New("cannot get detail category")
	}

	resp.ID = category.ID
	resp.Name = category.Name
	resp.Description = category.Description

	return resp, nil
}

// delete category by id
// update category by id
