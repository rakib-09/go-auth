package middlewares

//
//import (
//	"github.com/golang-jwt/jwt"
//	"github.com/labstack/echo/v4"
//	"go-auth/config"
//)
//
//func resolveJwtUser() echo.MiddlewareFunc {
//	return func(next echo.HandlerFunc) echo.HandlerFunc {
//		return func(c echo.Context) error {
//			user := c.Get("user").(*jwt.Token)
//			claims := user.Claims.(*config.JwtCustomClaims)
//			userId := claims.UID
//			user =
//			return next(c)
//		}
//	}
//}
