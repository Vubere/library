package middlewares

import (
	"net/http"
	"strconv"
	"strings"
	"time"
	"victorubere/library/jwt_manager"
	"victorubere/library/lib/helpers"
	"victorubere/library/lib/library_constants"
	"victorubere/library/models"
	"victorubere/library/services"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

func ValidateUserId(userService services.IUserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
			c.Abort()

			return
		}
		user, err := userService.GetUserById(userId)
		if err != nil {
			if err.Error() == "record not found" {
				c.JSON(http.StatusBadRequest, gin.H{"message": "user not found"})
				c.Abort()

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
				c.Abort()

				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"message": "error", "error": err.Error()})
			c.Abort()
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
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "error", "error": "something went wrong on the server"})
			c.Abort()

			return
		}
		if role != user.Role {
			c.JSON(http.StatusForbidden, gin.H{"message": "your role is not authorized to access this resource", "status": http.StatusForbidden})
			c.Abort()
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
			c.Abort()
			return
		}
		tokenString := AuthorizationSlice[1]
		claims, err := jwt_manager.VerifyToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized", "status": http.StatusUnauthorized})
			c.Abort()
			return
		}
		//check if the token is expired
		if claims["exp"].(float64) < float64(time.Now().Unix()) {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized: token expired", "status": http.StatusUnauthorized})
			c.Abort()
			return
		}
		user_id := claims["user_id"].(float64)
		user, err := userService.GetUserById(int(user_id))
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized", "status": http.StatusUnauthorized})
			c.Abort()
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
			c.Abort()

			return
		}
		user, ok := c.MustGet("User").(models.User)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "error", "error": "something went wrong on the server"})
			c.Abort()

			return
		}
		if !helpers.CompareHashedPasswordWithPlaintext(user.Password, input.Password) {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "invalid password", "status": http.StatusUnauthorized})
			c.Abort()

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
			c.Abort()

			return
		}
		user, ok := c.MustGet("User").(models.User)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "error", "error": "something went wrong on the server"})
			c.Abort()

			return
		}
		if user.Role == library_constants.ROLE_ADMIN {
			c.Next()
			return
		}
		if user.ID != uint(id) {
			c.JSON(http.StatusForbidden, gin.H{"message": "your role is not authorized to access this resource", "status": http.StatusForbidden})
			c.Abort()

			return
		}

		c.Next()
	}
}

func RateLimiter() gin.HandlerFunc {
	limiter := rate.NewLimiter(1, 3)
	return func(c *gin.Context) {

		if limiter.Allow() {
			c.Next()
		} else {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"message": "Limite exceed",
			})
		}

	}
}
