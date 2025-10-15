package controller

import (
	"net/http"
	"strconv"

	"github.com/Arroziqi/car-rental-technical-test-pharos.git/features/customer/application/usecase"
	custEntity "github.com/Arroziqi/car-rental-technical-test-pharos.git/features/customer/domain/entity"
	"github.com/gin-gonic/gin"
)

type CustomerController struct {
	uc *usecase.CustomerUsecase
}

func NewCustomerController(u *usecase.CustomerUsecase) *CustomerController {
	return &CustomerController{uc: u}
}

func (c *CustomerController) Create(ctx *gin.Context) {
	var req custEntity.Customer
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.uc.Create(ctx, &req); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, req)
}

func (c *CustomerController) List(ctx *gin.Context) {
	list, err := c.uc.List(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, list)
}

func (c *CustomerController) GetByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	item, err := c.uc.GetByID(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	ctx.JSON(http.StatusOK, item)
}

func (c *CustomerController) Update(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var req custEntity.Customer
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.ID = id
	if err := c.uc.Update(ctx, &req); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, req)
}

func (c *CustomerController) Delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := c.uc.Delete(ctx, id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusNoContent)
}
