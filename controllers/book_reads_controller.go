package controllers

import (
	"net/http"
	"strconv"
	"victorubere/library/lib/helpers"
	"victorubere/library/lib/structs"
	"victorubere/library/models"

	"github.com/gin-gonic/gin"
)

func (c *Controller) BookReadsController(rg *gin.RouterGroup) {
	bookReadsRoutes := rg.Group("/book-reads")
	{
		bookReadsRoutes.GET("", c.GetAllBookReads)
		bookReadsRoutes.POST("", c.CreateBookRead)
		bookReadsRoutes.GET("/:id", c.GetBookReadById)
		bookReadsRoutes.PUT("/:id", c.UpdateBookRead)
		bookReadsRoutes.DELETE("/:id", c.DeleteBookRead)
	}
}

func (c *Controller) GetAllBookReads(ctx *gin.Context) {
	var query structs.Query
	var bookReadQuery structs.BookReadQuery
	err := helpers.BindQuery(ctx, &query)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}
	err = helpers.BindModelQuery(ctx, &bookReadQuery)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	bookReads, count, err := c.bookReadService.GetAllBookReads(query, bookReadQuery)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error"})
		return
	}
	meta := helpers.GenerateMeta(count, query)
	ctx.JSON(http.StatusOK, gin.H{
		"book_reads":   bookReads,
		"status":  http.StatusOK,
		"message": "success",
		"meta":    meta,
	})
}

func (c *Controller) CreateBookRead(ctx *gin.Context) {
	var bookRead models.BookRead
	err := ctx.BindJSON(&bookRead)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request", "error": err.Error()})
		return
	}
	createdBookRead, err := c.bookReadService.CreateBookRead(bookRead, c.userService, c.bookService, c.visitationService)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error", "error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"book_read": createdBookRead, "status": http.StatusCreated, "message": "success"})
}

func (c *Controller) GetBookReadById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request", "error": err.Error()})
		return
	}
	bookRead, err := c.bookReadService.GetBookReadById(id)
	if err != nil {
		if err.Error() == "record not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"message": "record not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "failed to get book read", "error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"book_read": bookRead, "status": http.StatusOK, "message": "success"})
}

func (c *Controller) UpdateBookRead(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}	
	bookRead, err := c.bookReadService.GetBookReadById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error"})
		return
	}
	err = ctx.BindJSON(&bookRead)
	if err != nil {	
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}
	updatedBookRead, err := c.bookReadService.UpdateBookRead(bookRead, c.userService, c.bookService, c.visitationService)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"book_read": updatedBookRead,
		"status":  http.StatusOK,
		"message": "success",
	})
}

func (c *Controller) DeleteBookRead(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}
	err = c.bookReadService.DeleteBookRead(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "deleted", "status": http.StatusOK})
}