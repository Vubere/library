package controllers

import (
	"net/http"
	"victorubere/library/jwt_manager"
	"victorubere/library/lib/helpers"
	"victorubere/library/middlewares"
	"victorubere/library/models"

	"github.com/gin-gonic/gin"
)

func (c *Controller) AuthController(rg *gin.RouterGroup) {
	usersRoutes := rg.Group("/auth")
	{
		usersRoutes.POST("/login", c.LoginUser)
		usersRoutes.POST("/register", middlewares.ValidateUserEmail(c.userService), c.UserRegistration)
	}
}

func (c *Controller) LoginUser(ctx *gin.Context) {
	var input struct {
		Password string `json:"password"`
		Email    string `json:"email"`
	}
	err := ctx.BindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request", "error": err.Error()})
		return
	}
	
	user, err := c.userService.GetUserByEmail(input.Email)
	if err != nil {
		if err.Error() == "record not found" {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid username or password"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error", "error": err.Error()})
		return
	}
	if !helpers.CompareHashedPasswordWithPlaintext(user.Password, input.Password) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "invalid email or password", "status": http.StatusUnauthorized})
		return
	}
	token, err := jwt_manager.CreateToken(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error", "error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"user":    user,
		"token":   token,
		"message": "login successful",
		"status":  http.StatusOK,
	})
}

func (c *Controller) UserRegistration(ctx *gin.Context) {
	var user models.User
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request", "error": err.Error()})
		return
	}
	//check if email is already taken
	_, err = c.userService.GetUserByEmail(user.Email)
	if err.Error() != "record not found" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "email already taken", "error": err.Error()})
		return
	}
	//check if password is strong enough
	err = helpers.ValidatePasswordPlaintext(user.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request", "error": err.Error()})
		return
	}
	//create user
	user, err = c.userService.RegisterUser(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error", "error": err.Error()})
		return
	}
	token, err := jwt_manager.CreateToken(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error", "error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"user":    user,
		"token":   token,
		"message": "user created",
		"status":  http.StatusOK,
	})
}
