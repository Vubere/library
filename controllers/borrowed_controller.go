package controllers

import (
	"net/http"
	"strconv"
	"strings"
	"victorubere/library/lib/helpers"
	"victorubere/library/lib/structs"
	"victorubere/library/models"

	"github.com/gin-gonic/gin"
)

func (c *Controller) BorrowedController(rg *gin.RouterGroup) {
	borrowedsRoutes := rg.Group("/borroweds")
	{
		borrowedsRoutes.GET("", c.GetAllBorroweds)
		borrowedsRoutes.POST("", c.CreateBorrowed)
		borrowedsRoutes.GET("/:id", c.GetBorrowedById)
		borrowedsRoutes.PUT("/:id", c.UpdateBorrowed)
		borrowedsRoutes.DELETE("/:id", c.DeleteBorrowed)
	}
}

func (c *Controller) GetAllBorroweds(ctx *gin.Context) {
	var query structs.Query
	var borrowedQuery structs.BorrowedQuery
	err := helpers.BindQuery(ctx, &query)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request", "error": err.Error()})
		return
	}
	err = helpers.BindModelQuery(ctx, &borrowedQuery)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request", "error": err.Error()})
		return
	}

	borroweds, count, err := c.borrowedService.GetAllBorroweds(query, borrowedQuery)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error"})
		return
	}
	meta := helpers.GenerateMeta(count, query)
	ctx.JSON(http.StatusOK, gin.H{
		"borroweds":   borroweds,
		"status":  http.StatusOK,
		"message": "success",
		"meta":    meta,
	})
}

func (c *Controller) CreateBorrowed(ctx *gin.Context) {
	var borrowed models.Borrowed
	err := ctx.BindJSON(&borrowed)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}
	createdBorrowed, err := c.borrowedService.CreateBorrowed(borrowed, c.userService, c.bookService)
	if err != nil {
		errText := err.Error()
		if strings.Contains(errText, "not found") {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": errText, "status": http.StatusBadRequest})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error", "error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"borrowed": createdBorrowed})
}

func (c *Controller) GetBorrowedById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request", "error": err.Error()})
		return
	}
	borrowed, err := c.borrowedService.GetBorrowedById(id)
	if err != nil {
		if err.Error() == "record not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"message": "record not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error", "error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"borrowed": borrowed})
}

func (c *Controller) UpdateBorrowed(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}
	borrowed, err := c.borrowedService.GetBorrowedById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error"})
		return
	}
	err = ctx.BindJSON(&borrowed)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}
	updatedBorrowed, err := c.borrowedService.UpdateBorrowed(borrowed, c.userService, c.bookService)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"borrowed": updatedBorrowed, "message": "borrowed updated successfully", "status": http.StatusOK})
}

func (c *Controller) DeleteBorrowed(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request", "error": err.Error()})
		return
	}
	err = c.borrowedService.DeleteBorrowed(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error", "error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}
