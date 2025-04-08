package jwt_manager

import (
	"fmt"
	"os"
	"time"
	"victorubere/library/models"

	"github.com/golang-jwt/jwt/v5"
)

func getSecretKeyFromEnv() []byte {
	sk := os.Getenv("JWT_SECRET_KEY")
	if sk == "" {
		sk = "Secretkey"
	}
	return []byte(sk)
}

func CreateToken(user models.User) (string, error) {
	secretKey := getSecretKeyFromEnv()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"user_name": user.Name,
			"user_id":   user.ID,
			"role":      user.Role,
			"email":     user.Email,
			"exp":       time.Now().Add(time.Hour * 24).Unix(),
			"iss":       "https://victorubere.netlify.app",
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) (jwt.MapClaims, error) {
	secretKey := getSecretKeyFromEnv()
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("could not parse claims")
	}

	return claims, nil
}
