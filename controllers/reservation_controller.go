package controllers

import (
	"net/http"
	"strconv"
	"victorubere/library/lib/helpers"
	"victorubere/library/lib/structs"
	"victorubere/library/models"

	"github.com/gin-gonic/gin"
)

func (c *Controller) ReservationController(rg *gin.RouterGroup) {
	reservationsRoutes := rg.Group("/reservations")
	{
		reservationsRoutes.GET("", c.GetAllReservations)
		reservationsRoutes.POST("", c.CreateReservation)
		reservationsRoutes.GET("/:id", c.GetReservationById)
		reservationsRoutes.PUT("/:id", c.UpdateReservation)
		reservationsRoutes.DELETE("/:id", c.DeleteReservation)
	}
}

func (c *Controller) GetAllReservations(ctx *gin.Context) {
	var query structs.Query
	err := helpers.BindQuery(ctx, &query)
	if err != nil {
		return
	}


	reservations, err := c.reservationService.GetAllReservations(query)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"reservations": reservations})
}

func (c *Controller) CreateReservation(ctx *gin.Context) {
	var reservation models.Reservation
	err := ctx.BindJSON(&reservation)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}
	createdReservation, err := c.reservationService.CreateReservation(reservation)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"reservation": createdReservation})
}

func (c *Controller) GetReservationById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}
	reservation, err := c.reservationService.GetReservationById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"reservation": reservation})
}

func (c *Controller) UpdateReservation(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}
	reservation, err := c.reservationService.GetReservationById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error"})
		return
	}
	err = ctx.BindJSON(&reservation)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}
	updatedReservation, err := c.reservationService.UpdateReservation(reservation)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"reservation": updatedReservation})
}

func (c *Controller) DeleteReservation(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}
	err = c.reservationService.DeleteReservation(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}
