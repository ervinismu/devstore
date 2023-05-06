package service

import (
	"errors"
	"testing"

	"github.com/ervinismu/devstore/internal/app/mocks"
	"github.com/ervinismu/devstore/internal/app/model"
	"github.com/ervinismu/devstore/internal/app/schema"
	"github.com/ervinismu/devstore/internal/pkg/reason"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestProductService_BrowseAll(t *testing.T) {
	type TestCase struct {
		Name   string
		Given  []model.Category
		Expect int
		Error  error
	}

	cases := []TestCase{
		{
			Name: "when have 2 category data",
			Given: []model.Category{{
				ID:          1,
				Name:        "category 1",
				Description: "description category 1",
			}, {
				ID:          2,
				Name:        "category 2",
				Description: "description category 2",
			}},
			Expect: 2,
			Error:  nil,
		},
		{
			Name:   "when empty category data",
			Given:  []model.Category{},
			Expect: 0,
			Error:  nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockCategoryRepo := mocks.NewMockCategoryRepository(mockCtrl)
			dbSearch := model.BrowseProduct{}
			dbSearch.Page = 1
			dbSearch.PageSize = 2
			mockCategoryRepo.
				EXPECT().
				Browse(dbSearch).
				Return(tc.Given, tc.Error)

			req := &schema.BrowseCategoryReq{}
			req.Page = 1
			req.PageSize = 2
			categoryService := NewCategoryService(mockCategoryRepo)
			categories, err := categoryService.BrowseAll(req)
			total := len(categories)

			assert.Equal(t, tc.Expect, total)
			assert.NoError(t, err)
		})
	}
}

func TestProductService_GetByID(t *testing.T) {
	type TestCase struct {
		Name       string
		CategoryID string
		Given      model.Category
		Expect     int
		Error      error
	}

	cases := []TestCase{
		{
			Name:       "when category exist",
			CategoryID: "1",
			Given: model.Category{
				ID:          1,
				Name:        "category 1",
				Description: "description category 1",
			},
			Expect: 1,
			Error:  nil,
		},
		{
			Name:       "when category not exist",
			CategoryID: "2",
			Given:      model.Category{},
			Expect:     0,
			Error:      nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)

			mockCategoryRepo := mocks.NewMockCategoryRepository(mockCtrl)
			mockCategoryRepo.
				EXPECT().
				GetByID(tc.CategoryID).
				Return(tc.Given, tc.Error)

			mockCategoryRepo.
				EXPECT().
				GetByID(tc.CategoryID).
				Return(tc.Given, tc.Error)

			categoryService := NewCategoryService(mockCategoryRepo)
			category, _ := categoryService.GetByID(tc.CategoryID)

			assert.Equal(t, tc.Expect, category.ID)
		})
	}
}

func TestProductService_DeleteByID(t *testing.T) {
	t.Run("when category not exist", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)

		mockCategoryRepo := mocks.NewMockCategoryRepository(mockCtrl)
		mockCategoryRepo.
			EXPECT().
			GetByID("1").
			Return(model.Category{}, errors.New(reason.CategoryNotFound)).
			AnyTimes()
		mockCategoryRepo.
			EXPECT().
			DeleteByID("1").
			Return(nil).
			AnyTimes()

		categoryService := NewCategoryService(mockCategoryRepo)
		err := categoryService.DeleteByID("1")

		assert.EqualError(t, err, reason.CategoryNotFound)
	})
	t.Run("when category exist", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)

		mockCategoryRepo := mocks.NewMockCategoryRepository(mockCtrl)
		mockCategoryRepo.
			EXPECT().
			GetByID("1").
			Return(model.Category{}, nil).
			AnyTimes()
		mockCategoryRepo.
			EXPECT().
			DeleteByID("1").
			Return(nil).
			AnyTimes()

		categoryService := NewCategoryService(mockCategoryRepo)
		err := categoryService.DeleteByID("1")

		assert.Nil(t, err)
	})
	t.Run("when error delete category", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)

		mockCategoryRepo := mocks.NewMockCategoryRepository(mockCtrl)
		mockCategoryRepo.
			EXPECT().
			GetByID("1").
			Return(model.Category{}, nil).
			AnyTimes()
		mockCategoryRepo.
			EXPECT().
			DeleteByID("1").
			Return(errors.New(reason.CategoryCannotDelete)).
			AnyTimes()

		categoryService := NewCategoryService(mockCategoryRepo)
		err := categoryService.DeleteByID("1")

		assert.EqualError(t, err, reason.CategoryCannotDelete)
	})
}

func TestCategoryService_UpdateByID(t *testing.T) {
	t.Run("when category not exist", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)

		mockCategoryRepo := mocks.NewMockCategoryRepository(mockCtrl)
		mockCategoryRepo.
			EXPECT().
			GetByID("1").
			Return(model.Category{}, errors.New(reason.CategoryNotFound)).
			AnyTimes()
		mockCategoryRepo.
			EXPECT().
			Update("1").
			Return(nil).
			AnyTimes()

		categoryService := NewCategoryService(mockCategoryRepo)
		req := &schema.UpdateCategoryReq{}
		err := categoryService.UpdateByID("1", req)

		assert.EqualError(t, err, reason.CategoryNotFound)
	})
}
