package authentication

import (
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateToken(secret string, id uint, role uint) (string, error) {
	secretKey := []byte(secret)
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 24 * 365).Unix()

	tokenString, err := token.SignedString(secretKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseToken(secret string, tokenString string) (uint, uint, error) {
	secretKey := []byte(secret)
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		id := claims["id"].(float64)
		roleID := claims["role"].(float64)

		return uint(id), uint(roleID), nil
	} else {
		return 0, 0, err
	}
}
