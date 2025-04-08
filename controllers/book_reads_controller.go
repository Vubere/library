package controllers

import (
	"net/http"
	"strconv"
	"victorubere/library/lib/helpers"
	"victorubere/library/lib/library_contants"
	"victorubere/library/lib/structs"
	"victorubere/library/middlewares"
	"victorubere/library/models"

	"github.com/gin-gonic/gin"
)

func (c *Controller) BookReadssController(rg *gin.RouterGroup) {
	bookReadsRoutes := rg.Group("/book-reads")
	{
		bookReadsRoutes.GET("", middlewares.ValidateJWT(c.userService), c.GetAllBookReadss)
		bookReadsRoutes.POST("", middlewares.ValidateJWT(c.userService),  middlewares.ValidateUserRole(library_contants.ROLE_ADMIN, c.userService), c.CreateBookReads)
		bookReadsRoutes.GET("/:id", middlewares.ValidateJWT(c.userService), c.GetBookReadsById)
		bookReadsRoutes.PUT("/:id", middlewares.ValidateJWT(c.userService), middlewares.ValidateUserRole(library_contants.ROLE_ADMIN, c.userService), c.UpdateBookReads)
		bookReadsRoutes.DELETE("/:id", middlewares.ValidateJWT(c.userService), middlewares.ValidateUserRole(library_contants.ROLE_ADMIN, c.userService), c.DeleteBookReads)
	}
}

func (c *Controller) GetAllBookReadss(ctx *gin.Context) {
	var query structs.Query
	var bookReadQuery structs.BookReadsQuery
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

	bookReads, count, err := c.bookReadService.GetAllBookReadss(query, bookReadQuery)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error", "error": err.Error()})
		return
	}
	meta := helpers.GenerateMeta(count, query)
	ctx.JSON(http.StatusOK, gin.H{
		"book_reads": bookReads,
		"status":     http.StatusOK,
		"message":    "success",
		"meta":       meta,
	})
}

func (c *Controller) CreateBookReads(ctx *gin.Context) {
	var bookRead models.BookReads
	err := ctx.BindJSON(&bookRead)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request", "error": err.Error()})
		return
	}
	createdBookReads, err := c.bookReadService.CreateBookReads(bookRead, c.userService, c.bookService, c.visitationService)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error", "error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"book_read": createdBookReads, "status": http.StatusCreated, "message": "success"})
}

func (c *Controller) GetBookReadsById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request", "error": err.Error()})
		return
	}
	bookRead, err := c.bookReadService.GetBookReadsById(id)
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

func (c *Controller) UpdateBookReads(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}
	bookRead, err := c.bookReadService.GetBookReadsById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error"})
		return
	}
	err = ctx.BindJSON(&bookRead)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}
	updatedBookReads, err := c.bookReadService.UpdateBookReads(bookRead, c.userService, c.bookService, c.visitationService)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error", "error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"book_read": updatedBookReads,
		"status":    http.StatusOK,
		"message":   "success",
	})
}

func (c *Controller) DeleteBookReads(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}
	err = c.bookReadService.DeleteBookReads(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "deleted", "status": http.StatusOK})
}
