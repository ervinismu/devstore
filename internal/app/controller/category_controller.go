package controller

import (
	"net/http"

	"github.com/ervinismu/devstore/internal/app/schema"
	"github.com/ervinismu/devstore/internal/app/service"
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
		ctx.JSON(http.StatusUnprocessableEntity, gin.H {"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": resp})
}

func (cc *CategoryController) CreateCategory(ctx *gin.Context) {
	var req schema.CreateCategoryReq
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H {"message": err.Error()})
		return
	}

	err = cc.service.Create(req)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H {"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "success create category"})
}

// get detail category
func (cc *CategoryController) DetailCategory(ctx *gin.Context) {
	id, _ := ctx.Params.Get("id")
	resp, err := cc.service.GetByID(id)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H {"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H { "data": resp })
}

// update article by id
func (cc *CategoryController) UpdateCategory(ctx *gin.Context) {
	id, _ := ctx.Params.Get("id")
	var req schema.UpdateCategoryReq

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H {"message": err.Error()})
		return
	}

	err = cc.service.UpdateByID(id, req)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H {"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "success update category"})
}

// delete article by id
func (cc *CategoryController) DeleteCategory(ctx *gin.Context) {
	id, _ := ctx.Params.Get("id")

	err := cc.service.DeleteByID(id)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H {"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H { "message": "success delete category" })
}
