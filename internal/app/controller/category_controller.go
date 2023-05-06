package controller

import (
	"net/http"

	"github.com/ervinismu/devstore/internal/app/schema"
	"github.com/ervinismu/devstore/internal/pkg/handler"
	"github.com/gin-gonic/gin"
)

type CategoryService interface {
	Create(req *schema.CreateCategoryReq) error
	BrowseAll(req *schema.BrowseCategoryReq) ([]schema.GetCategoryResp, error)
	GetByID(id string) (schema.GetCategoryResp, error)
	UpdateByID(id string, req *schema.UpdateCategoryReq) error
	DeleteByID(id string) error
}

type CategoryController struct {
	service CategoryService
}

func NewCategoryController(service CategoryService) *CategoryController {
	return &CategoryController{service: service}
}

// create category
func (cc *CategoryController) BrowseCategory(ctx *gin.Context) {
	req := &schema.BrowseCategoryReq{}
	req.Page = ctx.GetInt("page")
	req.PageSize = ctx.GetInt("page_size")

	resp, err := cc.service.BrowseAll(req)
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

	handler.ResponseSuccess(ctx, http.StatusOK, "success create category", nil)
}

// get detail category
func (cc *CategoryController) DetailCategory(ctx *gin.Context) {
	id, _ := ctx.Params.Get("id")
	resp, err := cc.service.GetByID(id)
	if err != nil {
		handler.ResponseError(ctx, http.StatusOK, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "", resp)
}

// update category by id
func (cc *CategoryController) UpdateCategory(ctx *gin.Context) {
	id, _ := ctx.Params.Get("id")
	req := &schema.UpdateCategoryReq{}
	if handler.BindAndCheck(ctx, req) {
		return
	}

	err := cc.service.UpdateByID(id, req)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success update category", nil)
}

// delete category by id
func (cc *CategoryController) DeleteCategory(ctx *gin.Context) {
	id, _ := ctx.Params.Get("id")

	err := cc.service.DeleteByID(id)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success delete category", nil)
}
