package controllers

import (
	"net/http"
	"strconv"
	"victorubere/library/lib/helpers"
	"victorubere/library/lib/structs"
	"victorubere/library/models"

	"github.com/gin-gonic/gin"
)

func (c *Controller) BookController(rg *gin.RouterGroup) {
	booksRoutes := rg.Group("/books")
	{
		booksRoutes.GET("", c.GetAllBooks)
		booksRoutes.POST("", c.CreateBook)
		booksRoutes.GET("/:id", c.GetBookById)
		booksRoutes.PUT("/:id", c.UpdateBook)
		booksRoutes.DELETE("/:id", c.DeleteBook)
	}
}

func (c *Controller) GetAllBooks(ctx *gin.Context) {
	var query structs.Query
	var bookQuery structs.BookQuery
	err := helpers.BindQuery(ctx, &query)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}
	err = helpers.BindModelQuery(ctx, &bookQuery)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	books, count, err := c.bookService.GetAllBooks(query, bookQuery)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error"})
		return
	}
	meta := helpers.GenerateMeta(count, query)
	ctx.JSON(http.StatusOK, gin.H{
		"books":   books,
		"meta":    meta,
		"status":  http.StatusOK,
		"message": "success",
	})
}

func (c *Controller) CreateBook(ctx *gin.Context) {
	var book models.Book
	err := ctx.BindJSON(&book)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request",
			"error":   err.Error(),
		})
		return
	}
	createdBook, err := c.bookService.CreateBook(book)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "error",
			"error":   err.Error(),
			"status":  http.StatusInternalServerError,
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"book":    createdBook,
		"status":  http.StatusCreated,
		"message": "success",
	})
}

func (c *Controller) GetBookById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
		return
	}
	book, err := c.bookService.GetBookById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "error",
			"status":  http.StatusNotFound,
			"error":   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"book":    book,
		"status":  http.StatusOK,
		"message": "success",
	})
}

func (c *Controller) UpdateBook(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}
	book, err := c.bookService.GetBookById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
			"status":  http.StatusNotFound,
		})
		return
	}
	err = ctx.BindJSON(&book)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request",
			"error":   err.Error(),
			"status":  http.StatusBadRequest,
		})
		return
	}
	updatedBook, err := c.bookService.UpdateBook(book)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "error",
			"error":   err.Error(),
			"status":  http.StatusInternalServerError,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"book":    updatedBook,
		"status":  http.StatusOK,
		"message": "Book  updated!",
	})
}

func (c *Controller) DeleteBook(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}
	err = c.bookService.DeleteBook(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "book deleted", "status": http.StatusOK})
}
