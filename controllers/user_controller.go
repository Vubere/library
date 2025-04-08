package controllers

import (
	"net/http"
	"strconv"
	"victorubere/library/lib/helpers"
	"victorubere/library/lib/library_constants"
	"victorubere/library/lib/structs"
	"victorubere/library/middlewares"
	"victorubere/library/models"

	"github.com/gin-gonic/gin"
)

func (c *Controller) UserController(rg *gin.RouterGroup) {
	usersRoutes := rg.Group("/users")
	{
		usersRoutes.GET("", middlewares.ValidateJWT(c.userService), middlewares.ValidateUserRole(library_constants.ROLE_ADMIN, c.userService), c.GetAllUsers)
		usersRoutes.POST("", middlewares.ValidateJWT(c.userService), middlewares.ValidateUserRole(library_constants.ROLE_ADMIN, c.userService), c.RegisterUser)
		usersRoutes.GET("/:id", middlewares.ValidateJWT(c.userService), c.GetUserById)
		usersRoutes.PUT("/:id", middlewares.ValidateJWT(c.userService), middlewares.ValidateUserId(c.userService), middlewares.ConfirmThatUserHasID(c.userService), c.UpdateUser)
		usersRoutes.DELETE("/:id",  middlewares.ValidateJWT(c.userService), middlewares.ValidateUserId(c.userService), middlewares.ConfirmThatUserHasID(c.userService), c.DeleteUser)
		usersRoutes.GET("/summary/:id",  middlewares.ValidateJWT(c.userService), middlewares.ValidateUserId(c.userService), c.UserSummary)
	}
}

func (c *Controller) GetAllUsers(ctx *gin.Context) {
	var query structs.Query
	var userQuery structs.UserQuery
	err := helpers.BindQuery(ctx, &query)
	if err != nil {
		return
	}
	err = helpers.BindModelQuery(ctx, &userQuery)
	if err != nil {
		return
	}

	users, count, err := c.userService.GetAllUsers(query, userQuery)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error", "error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"users":   users,
		"status":  http.StatusOK,
		"message": "success",
		"meta":    helpers.GenerateMeta(count, query),
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
	var returnedUser models.UserDTO
	returnedUser.ID = createdUser.ID
	returnedUser.CreatedAt = &createdUser.CreatedAt
	returnedUser.UpdatedAt = &createdUser.UpdatedAt
	returnedUser.Name = createdUser.Name
	returnedUser.Email = createdUser.Email
	returnedUser.PhoneNumber = createdUser.PhoneNumber
	returnedUser.Address = createdUser.Address
	returnedUser.Gender = createdUser.Gender
	returnedUser.Role = createdUser.Role
	ctx.JSON(http.StatusOK, gin.H{
		"user":    returnedUser,
		"token":   "",
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
		if err.Error() == "record not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "error",
			"error":   err.Error(),
		})
		return
	}
	var returnedUser models.UserDTO
	returnedUser.ID = user.ID
	returnedUser.CreatedAt = &user.CreatedAt
	returnedUser.UpdatedAt = &user.UpdatedAt
	returnedUser.Name = user.Name
	returnedUser.Email = user.Email
	returnedUser.PhoneNumber = user.PhoneNumber
	returnedUser.Address = user.Address
	returnedUser.Gender = user.Gender
	returnedUser.Role = user.Role
	ctx.JSON(http.StatusOK, gin.H{"user": returnedUser, "status": http.StatusOK, "message": "success"})
}

func (c *Controller) UpdateUser(ctx *gin.Context) {
	user, ok := ctx.MustGet("User").(models.User)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error", "error": "something went wrong on the server"})
		return
	}
	if user.Role != library_constants.ROLE_USER && user.Role != library_constants.ROLE_ADMIN {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid user role", "status": http.StatusBadRequest})
		return
	}
	if user.Password != "" {
		user.Password = ""
	}
	updatedUser, err := c.userService.UpdateUser(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "failed to update user"})
		return
	}
	var returnedUser models.UserDTO
	returnedUser.ID = updatedUser.ID
	returnedUser.CreatedAt = &updatedUser.CreatedAt
	returnedUser.UpdatedAt = &updatedUser.UpdatedAt
	returnedUser.Name = updatedUser.Name
	returnedUser.Email = updatedUser.Email
	returnedUser.PhoneNumber = updatedUser.PhoneNumber
	returnedUser.Address = updatedUser.Address
	returnedUser.Gender = updatedUser.Gender
	returnedUser.Role = updatedUser.Role
	ctx.JSON(http.StatusOK, gin.H{
		"user":    returnedUser,
		"status":  http.StatusOK,
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

func (c *Controller) UserSummary(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}
	_, err = c.userService.GetUserById(id)
	if err != nil {
		if err.Error() == "record not found" {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "user not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error", "error": err.Error()})
		return
	}
	userSummary, err := c.userService.GetUserSummary(id, c.visitationService, c.borrowedService, c.bookReadService)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error", "error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"user": userSummary, "status": http.StatusOK, "message": "success"})
}
