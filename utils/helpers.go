package utils

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"go-auth/config"
)

func GetUserIdFromJwt(c echo.Context) uint {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*config.JwtCustomClaims)
	return claims.UID
}
