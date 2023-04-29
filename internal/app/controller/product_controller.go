package controller

import (
	"net/http"

	"github.com/ervinismu/devstore/internal/app/schema"
	"github.com/ervinismu/devstore/internal/pkg/handler"
	"github.com/gin-gonic/gin"
)

type ProductService interface {
	Create(req *schema.CreateProductReq) error
	BrowseAll() ([]schema.BrowseProductResp, error)
	GetByID(id string) (schema.DetailProductResp, error)
	UpdateByID(id string, req *schema.UpdateProductReq) error
	DeleteByID(id string) error
}

type ProductController struct {
	service ProductService
}

func NewProductController(service ProductService) *ProductController {
	return &ProductController{service: service}
}

// browse product
func (cc *ProductController) BrowseProduct(ctx *gin.Context) {
	resp, err := cc.service.BrowseAll()
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "", resp)
}

// create product
func (cc *ProductController) CreateProduct(ctx *gin.Context) {
	req := &schema.CreateProductReq{}
	if handler.BindAndCheck(ctx, req) {
		return
	}

	err := cc.service.Create(req)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success create product", nil)
}

// get detail product
func (cc *ProductController) DetailProduct(ctx *gin.Context) {
	id, _ := ctx.Params.Get("id")
	resp, err := cc.service.GetByID(id)
	if err != nil {
		handler.ResponseError(ctx, http.StatusOK, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "", resp)
}

// update product by id
func (cc *ProductController) UpdateProduct(ctx *gin.Context) {
	id, _ := ctx.Params.Get("id")
	req := &schema.UpdateProductReq{}
	if handler.BindAndCheck(ctx, req) {
		return
	}

	err := cc.service.UpdateByID(id, req)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success update product", nil)
}

// delete product by id
func (cc *ProductController) DeleteProduct(ctx *gin.Context) {
	id, _ := ctx.Params.Get("id")

	err := cc.service.DeleteByID(id)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success delete product", nil)
}
