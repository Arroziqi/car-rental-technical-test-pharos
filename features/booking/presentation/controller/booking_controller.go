package controller

import (
	"net/http"
	"strconv"

	"github.com/Arroziqi/car-rental-technical-test-pharos.git/features/booking/application/usecase"
	bookEntity "github.com/Arroziqi/car-rental-technical-test-pharos.git/features/booking/domain/entity"
	"github.com/gin-gonic/gin"
)

type BookingController struct {
	uc *usecase.BookingUsecase
}

func NewBookingController(u *usecase.BookingUsecase) *BookingController {
	return &BookingController{uc: u}
}

func (c *BookingController) Create(ctx *gin.Context) {
	var req bookEntity.Booking
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.uc.Create(ctx, &req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, req)
}

func (c *BookingController) List(ctx *gin.Context) {
	list, err := c.uc.List(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, list)
}

func (c *BookingController) GetByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	item, err := c.uc.GetByID(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	ctx.JSON(http.StatusOK, item)
}

func (c *BookingController) Update(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var req bookEntity.Booking
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

func (c *BookingController) Delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := c.uc.Delete(ctx, id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusNoContent)
}
