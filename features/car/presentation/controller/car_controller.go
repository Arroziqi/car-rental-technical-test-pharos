package controller

import (
	"net/http"
	"strconv"

	"github.com/Arroziqi/car-rental-technical-test-pharos.git/features/car/application/usecase"
	carEntity "github.com/Arroziqi/car-rental-technical-test-pharos.git/features/car/domain/entity"
	"github.com/gin-gonic/gin"
)

type CarController struct {
	uc *usecase.CarUsecase
}

func NewCarController(u *usecase.CarUsecase) *CarController {
	return &CarController{uc: u}
}

func (c *CarController) Create(ctx *gin.Context) {
	var req carEntity.Car
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

func (c *CarController) List(ctx *gin.Context) {
	list, err := c.uc.List(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, list)
}

func (c *CarController) GetByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	item, err := c.uc.GetByID(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	ctx.JSON(http.StatusOK, item)
}

func (c *CarController) Update(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var req carEntity.Car
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

func (c *CarController) Delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := c.uc.Delete(ctx, id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusNoContent)
}
