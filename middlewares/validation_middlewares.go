package middlewares

import (
	"net/http"
	"strconv"
	"strings"
	"victorubere/library/jwt_manager"
	"victorubere/library/lib/helpers"
	"victorubere/library/lib/library_contants"
	"victorubere/library/models"
	"victorubere/library/services"

	"github.com/gin-gonic/gin"
)

func ValidateUserId(userService services.IUserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
			return
		}
		user, err := userService.GetUserById(userId)
		if err != nil {
			if err.Error() == "record not found" {
				c.JSON(http.StatusBadRequest, gin.H{"message": "user not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"message": "error", "error": err.Error()})
			return
		}
		c.Set("User", user)
		c.Next()
	}
}

func ValidateUserEmail(userService services.IUserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		email := c.Param("email")
		user, err := userService.GetUserByEmail(email)
		if err != nil {
			if err.Error() == "record not found" {
				c.JSON(http.StatusBadRequest, gin.H{"message": "user not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"message": "error", "error": err.Error()})
			return
		}
		c.Set("User", user)
		c.Next()
	}
}

// ValidateUserRole should only be used after a middleware that inserts the user in the context
func ValidateUserRole(role string, userService services.IUserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, ok := c.MustGet("User").(models.User)
		if ok {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "error", "error": "something went wrong on the server"})
			return
		}
		if role != user.Role {
			c.JSON(http.StatusForbidden, gin.H{"message": "your role is not authorized to access this resource", "status": http.StatusForbidden})
			return
		}
		c.Next()
	}
}

// ValidateJWT should be used only for protected Routes
func ValidateJWT(userService services.IUserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		Authorization := c.Request.Header.Get("Authorization")
		AuthorizationSlice := strings.Split(Authorization, " ")
		if len(AuthorizationSlice) != 2 {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized", "status": http.StatusUnauthorized})
			return
		}
		tokenString := AuthorizationSlice[1]
		claims, err := jwt_manager.VerifyToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized", "status": http.StatusUnauthorized})
			return
		}
		user_id := claims["user_id"].(int)
		user, err := userService.GetUserById(user_id)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized", "status": http.StatusUnauthorized})
			return
		}
		c.Set("User", user)
		c.Next()
	}
}

func ValidateUserPassword(userService services.IUserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input struct {
			Password string `json:"password"`
		}
		err := c.ShouldBindJSON(&input)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request", "error": err.Error()})
			return
		}
		user, ok := c.MustGet("User").(models.User)
		if ok {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "error", "error": "something went wrong on the server"})
			return
		}
		if !helpers.CompareHashedPasswordWithPlaintext(user.Password, input.Password) {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "invalid password", "status": http.StatusUnauthorized})
			return
		}
		c.Next()
	}
}

func ConfirmThatUserHasID(userService services.IUserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
			return
		}
		user, ok := c.MustGet("User").(models.User)
		if ok {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "error", "error": "something went wrong on the server"})
			return
		}
		if user.Role == library_contants.ROLE_ADMIN {
			c.Next()
			return
		}
		if user.ID != uint(id) {
			c.JSON(http.StatusForbidden, gin.H{"message": "your role is not authorized to access this resource", "status": http.StatusForbidden})
			return
		}

		c.Next()
	}
}
