package middleware

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4/middleware"
	"time"
)

func CreateToken() string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Minute * 60).Unix()
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "Failed"
	}
	return t
}

var Auth = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: []byte("secret"),
})
