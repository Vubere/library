package controllers

import (
	"log"
	"net/http"
	"victorubere/library/jwt_manager"
	"victorubere/library/lib/helpers"
	"victorubere/library/lib/library_constants"
	"victorubere/library/models"

	"github.com/gin-gonic/gin"
)

func (c *Controller) AuthController(rg *gin.RouterGroup) {
	usersRoutes := rg.Group("/auth")
	{
		usersRoutes.POST("/login", c.LoginUser)
		usersRoutes.POST("/register", c.UserRegistration)
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
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid email or password", "status": http.StatusBadRequest})
		return
	}
	if user.Role != library_constants.ROLE_USER && user.Role != library_constants.ROLE_ADMIN {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid user role", "status": http.StatusBadRequest})
		return
	}
	token, err := jwt_manager.CreateToken(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error", "error": err.Error()})
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
	ctx.JSON(http.StatusOK, gin.H{
		"user":    returnedUser,
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
	fetchedUser, err := c.userService.GetUserByEmail(user.Email)
	if err != nil {
		if err.Error() != "record not found" {
			log.Printf("error: %v", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error"})
			ctx.Abort()
			return
		}
	}
	if fetchedUser.ID != 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "this email already has a account with us", "status": http.StatusBadRequest})
		ctx.Abort()
		return
	}
	//check if password is strong enough
	err = helpers.ValidatePasswordPlaintext(user.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request", "error": err.Error()})
		return
	}
	user.Password, err = helpers.EncryptPasswordFromPlaintext(user.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error", "error": err.Error()})
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
	ctx.JSON(http.StatusCreated, gin.H{
		"user":    returnedUser,
		"token":   token,
		"message": "user created",
		"status":  http.StatusCreated,
	})
}
