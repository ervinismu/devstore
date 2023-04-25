package service

import (
	"testing"

	"github.com/ervinismu/devstore/internal/app/mocks"
	"github.com/ervinismu/devstore/internal/app/model"
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
			mockCategoryRepo.EXPECT().Browse().Return(tc.Given, tc.Error)

			categoryService := NewCategoryService(mockCategoryRepo)
			categories, err := categoryService.BrowseAll()
			total := len(categories)

			assert.Equal(t, total, tc.Expect)
			assert.NoError(t, err)
		})
	}
}
