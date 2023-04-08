package service

import (
	"errors"
	"fmt"

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

// create category
func (cs *CategoryService) Create(req schema.CreateCategoryReq) error {
	var insertData model.Category

	insertData.Name = req.Name
	insertData.Description = req.Description

	err := cs.repo.Create(insertData)
	if err != nil {
		fmt.Println(err)
		return errors.New("cannot create category")
	}

	return nil
}

// get list category
func (cs *CategoryService) BrowseAll() ([]schema.GetCategoryResp, error) {
	var resp []schema.GetCategoryResp

	categories, err := cs.repo.Browse()
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("cannot get categories")
	}

	for _, value := range categories {
		var respData schema.GetCategoryResp
		respData.ID = value.ID
		respData.Name = value.Name
		respData.Description = value.Description
		resp = append(resp, respData)
	}

	return resp, nil
}

// get detail category
func (cs *CategoryService) GetByID(id string) (schema.GetCategoryResp, error) {
	var resp schema.GetCategoryResp

	category, err := cs.repo.GetByID(id)
	if err != nil {
		fmt.Println(err)
		return resp, errors.New("cannot get detail category")
	}

	resp.ID = category.ID
	resp.Name = category.Name
	resp.Description = category.Description

	return resp, nil
}

// update article by id
func (cs *CategoryService) UpdateByID(id string, req schema.UpdateCategoryReq) (error) {

	var updateData model.Category

	updateData.Name = req.Name
	updateData.Description = req.Description

	err := cs.repo.UpdateByID(id, updateData)
	if err != nil {
		fmt.Println(err)
		return errors.New("cannot update category")
	}

	return nil
}

// delete article by id
func (cs *CategoryService) DeleteByID(id string) (error) {

	err := cs.repo.DeleteByID(id)
	if err != nil {
		fmt.Println(err)
		return errors.New("cannot delete category")
	}

	return nil
}
