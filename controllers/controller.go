package controllers

import (
	"net/http"

	"victorubere/library/services"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	userService        services.IUserService
	bookService        services.IBookService
	visitationService  services.IVisitationService
	reservationService services.IReservationService
}

func NewController(userService services.IUserService, bookService services.IBookService, visitationService services.IVisitationService, reservationService services.IReservationService) *Controller {
	return &Controller{
		userService:        userService,
		bookService:        bookService,
		visitationService:  visitationService,
		reservationService: reservationService,
	}
}

func (c *Controller) InitializeRoutes() (*gin.Engine, error) {
	router := gin.Default()

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	v1 := router.Group("/api/v1")
	c.UserController(v1)
	c.BookController(v1)

	return router, nil
}
