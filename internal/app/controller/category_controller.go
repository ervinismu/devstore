package controller

import (
	"net/http"

	"github.com/ervinismu/devstore/internal/app/schema"
	"github.com/ervinismu/devstore/internal/app/service"
	"github.com/ervinismu/devstore/internal/pkg/handler"
	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	service service.ICategoryService
}

func NewCategoryController(service service.ICategoryService) *CategoryController {
	return &CategoryController{service: service}
}

// create category
func (cc *CategoryController) BrowseCategory(ctx *gin.Context) {
	resp, err := cc.service.BrowseAll()
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "", resp)
}

func (cc *CategoryController) CreateCategory(ctx *gin.Context) {
	req := &schema.CreateCategoryReq{}
	if handler.BindAndCheck(ctx, req) {
		return
	}

	err := cc.service.Create(req)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusCreated, "success create category", nil)
}

// get detail category
func (cc *CategoryController) DetailCategory(ctx *gin.Context) {
	id, _ := ctx.Params.Get("id")
	resp, err := cc.service.GetByID(id)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusCreated, "", resp)
}

// update article by id
func (cc *CategoryController) UpdateCategory(ctx *gin.Context) {
	req := &schema.UpdateCategoryReq{}
	if handler.BindAndCheck(ctx, req) {
		return
	}

	id, _ := ctx.Params.Get("id")
	err := cc.service.UpdateByID(id, req)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusCreated, "success update category", nil)
}

// delete article by id
func (cc *CategoryController) DeleteCategory(ctx *gin.Context) {
	id, _ := ctx.Params.Get("id")

	err := cc.service.DeleteByID(id)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusCreated, "success delete category", nil)
}
