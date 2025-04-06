package controllers

import (
	"net/http"
	"strconv"
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
	err := helpers.BindQuery(ctx, &query)
	if err != nil {
		return
	}

	borroweds, err := c.borrowedService.GetAllBorroweds(query)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"borroweds": borroweds})
}

func (c *Controller) CreateBorrowed(ctx *gin.Context) {
	var borrowed models.Borrowed
	err := ctx.BindJSON(&borrowed)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}
	createdBorrowed, err := c.borrowedService.CreateBorrowed(borrowed)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"borrowed": createdBorrowed})
}

func (c *Controller) GetBorrowedById(ctx *gin.Context) {
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
	updatedBorrowed, err := c.borrowedService.UpdateBorrowed(borrowed)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"borrowed": updatedBorrowed})
}

func (c *Controller) DeleteBorrowed(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}
	err = c.borrowedService.DeleteBorrowed(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}
