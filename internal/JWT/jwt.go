package JWT

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"
	"os"
	"time"
)

var jwtKey = []byte(os.Getenv("JWT_KEY"))

func GenerateToken(id string, tokenDuration time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":      id,
		"expires": tokenDuration,
		"issued":  time.Now().Unix(),
	})
	tokenString, err := token.SignedString(jwtKey)
	return tokenString, err
}

func ParseToken(tokenString string) (jwt.MapClaims, error) {
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("wrong token")
		}
		return jwtKey, nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("token is not valid")
}
