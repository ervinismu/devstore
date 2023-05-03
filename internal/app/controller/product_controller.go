package controller

import (
	"net/http"

	"github.com/ervinismu/devstore/internal/app/schema"
	"github.com/ervinismu/devstore/internal/app/service"
	"github.com/ervinismu/devstore/internal/pkg/handler"
	"github.com/ervinismu/devstore/internal/pkg/reason"
	"github.com/gin-gonic/gin"
)

type ProductController struct {
	service service.IProductService
}

func NewProductController(service service.IProductService) *ProductController {
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

	file, _ := req.Image.Open()
	defer file.Close()

	// validate file
	buff := make([]byte, 512)
	if _, err := file.Read(buff); err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, reason.InvalidImageFormat)
		return
	}

	filetype := http.DetectContentType(buff)
	if filetype != "image/jpeg" && filetype != "image/png" {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, reason.InvalidImageFormat)
		return
	}

	// ctx.Request.Body = http.MaxBytesReader(ctx.Writer, ctx.Request.Body, 1024)
	// if err := ctx.Request.ParseMultipartForm(1024); err != nil {
	// 	handler.ResponseError(ctx, http.StatusUnprocessableEntity, reason.InvalidImageFormat)
	// 	return
	// }

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
