package controllers

import (
	"net/http"
	"strconv"
	"victorubere/library/lib/helpers"
	"victorubere/library/lib/structs"
	"victorubere/library/models"

	"github.com/gin-gonic/gin"
)

func (c *Controller) VisitationController(rg *gin.RouterGroup) {
	visitationsRoutes := rg.Group("/visitations")
	{
		visitationsRoutes.GET("", c.GetAllVisitations)
		visitationsRoutes.POST("", c.CreateVisitation)
		visitationsRoutes.GET("/:id", c.GetVisitationById)
		visitationsRoutes.PUT("/:id", c.UpdateVisitation)
		visitationsRoutes.DELETE("/:id", c.DeleteVisitation)
	}
}

func (c *Controller) GetAllVisitations(ctx *gin.Context) {
	var query structs.Query
	var visitationQuery structs.VisitationQuery
	err := helpers.BindQuery(ctx, &query)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request", "error": err.Error()})
		return
	}
	err = helpers.BindModelQuery(ctx, &visitationQuery)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request", "error": err.Error()})
		return
	}

	visitations, count, err := c.visitationService.GetAllVisitations(query, visitationQuery)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "error",
			"error":   err.Error(),
		})
		return
	}
	meta := helpers.GenerateMeta(count, query)
	ctx.JSON(http.StatusOK, gin.H{
		"visitations": visitations,
		"status":      http.StatusOK,
		"message":     "success",
		"meta":        meta,
	})
}

func (c *Controller) CreateVisitation(ctx *gin.Context) {
	var visitation models.Visitations
	err := ctx.BindJSON(&visitation)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request", "error": err.Error()})
		return
	}
	createdVisitation, err := c.visitationService.CreateVisitation(visitation, c.userService)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error", "error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"visitation": createdVisitation})
}

func (c *Controller) GetVisitationById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}
	visitation, err := c.visitationService.GetVisitationById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error", "error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"visitation": visitation, "status": http.StatusCreated, "message": "success"})
}

func (c *Controller) UpdateVisitation(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}
	visitation, err := c.visitationService.GetVisitationById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	err = ctx.BindJSON(&visitation)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}
	updatedVisitation, err := c.visitationService.UpdateVisitation(visitation, c.userService)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error", "error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"visitation": updatedVisitation, "status": http.StatusOK, "message": "success"})
}

func (c *Controller) DeleteVisitation(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}
	err = c.visitationService.DeleteVisitation(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}
