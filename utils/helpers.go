package utils

import (
	"encoding/json"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"go-auth/config"
)

func GetUserIdFromJwt(c echo.Context) uint {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*config.JwtCustomClaims)
	return claims.UID
}

func MapStruct(input interface{}, output interface{}) error {
	if stringify, err := json.Marshal(input); err == nil {
		return json.Unmarshal(stringify, &output)
	} else {
		return err
	}
}
