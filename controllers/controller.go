package controllers

import (
	"net/http"

	"victorubere/library/services"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	userService       services.IUserService
	bookService       services.IBookService
	visitationService services.IVisitationService
	borrowedService   services.IBorrowedService
	bookReadService   services.IBookReadsService
}

func NewController(userService services.IUserService, bookService services.IBookService, visitationService services.IVisitationService, borrowedService services.IBorrowedService, bookReadService services.IBookReadsService) *Controller {
	return &Controller{
		userService:       userService,
		bookService:       bookService,
		visitationService: visitationService,
		borrowedService:   borrowedService,
		bookReadService:   bookReadService,
	}
}

func (c *Controller) InitializeRoutes() (*gin.Engine, error) {
	router := gin.Default()

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	//write endpoints to handle all options request
	router.OPTIONS("/*path", func(ctx *gin.Context) {
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		ctx.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
		ctx.Header("Access-Control-Allow-Credentials", "true")
		ctx.AbortWithStatus(http.StatusNoContent)
	})

	v1 := router.Group("/api/v1")
	c.UserController(v1)
	c.BookController(v1)
	c.VisitationController(v1)
	c.BorrowedController(v1)
	c.BookReadssController(v1)

	return router, nil
}
