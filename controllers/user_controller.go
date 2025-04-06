package controllers

import (
	"net/http"
	"strconv"
	"victorubere/library/lib/helpers"
	"victorubere/library/lib/structs"
	"victorubere/library/models"

	"github.com/gin-gonic/gin"
)

func (c *Controller) UserController(rg *gin.RouterGroup) {
	usersRoutes := rg.Group("/users")
	{
		usersRoutes.GET("", c.GetAllUsers)
		usersRoutes.POST("", c.RegisterUser)
		usersRoutes.GET("/:id", c.GetUserById)
		usersRoutes.PUT("/:id", c.UpdateUser)
		usersRoutes.DELETE("/:id", c.DeleteUser)
	}
}

func (c *Controller) GetAllUsers(ctx *gin.Context) {
	var query structs.Query
	var userQuery structs.UserQuery
	err := helpers.BindQuery(ctx, &query)
	if err != nil {
		return
	}
	err = helpers.BindUserQuery(ctx, &userQuery)
	if err != nil {
		return
	}

	users, count, err := c.userService.GetAllUsers(query, userQuery)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error", "error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"users": users,
		"status": http.StatusOK,
		"message": "success",
		"meta": helpers.GenerateMeta(count, query),
	})
}

func (c *Controller) RegisterUser(ctx *gin.Context) {
	var user models.User
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request", "error": err.Error()})
		return
	}
	createdUser, err := c.userService.RegisterUser(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error", "error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"user":    createdUser,
		"token":  "",
		"message": "user created",
		"status":  http.StatusOK,
	})
}

func (c *Controller) GetUserById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}
	user, err := c.userService.GetUserById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "error",
			"error":   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"user": user})
}

func (c *Controller) UpdateUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}
	user, err := c.userService.GetUserById(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"message": "failed",
			"error": err.Error(),
		})
		return
	}
	
	err = ctx.BindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}
	updatedUser, err := c.userService.UpdateUser(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "failed to update user"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"user": updatedUser,
		"status": http.StatusOK,
		"message": "success",
	})
}

func (c *Controller) DeleteUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}
	err = c.userService.DeleteUser(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "user deleted", "status": http.StatusOK})
}
