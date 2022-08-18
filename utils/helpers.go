package utils

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"go-auth/config"
)

func GetUserIdFromJwt(c echo.Context) int {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*config.JwtCustomClaims)
	return int(claims.UID)
}
