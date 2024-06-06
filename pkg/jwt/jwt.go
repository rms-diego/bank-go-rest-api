package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rms-diego/bank-go-rest-api/models"
	"github.com/rms-diego/bank-go-rest-api/pkg/config"
)

func CreateToken(user models.User) (string, error) {
	env := config.GetEnvironmentVariables()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":       user.Id,
			"name":     user.Name,
			"lastName": user.LastName,
			"email":    user.Email,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(env.JwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
