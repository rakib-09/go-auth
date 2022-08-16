package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go-auth/config"
	c "go-auth/controllers"
	m "go-auth/middlewares"
)

type Routes struct {
	echo           *echo.Echo
	UserController *c.UserController
	AuthController *c.AuthController
}

func New(e *echo.Echo, uc *c.UserController, ac *c.AuthController) *Routes {
	return &Routes{
		echo:           e,
		UserController: uc,
		AuthController: ac,
	}
}

func (r *Routes) Init() {
	e := r.echo
	m.Init(r.echo)
	g := e.Group("/api/v1")

	// public
	g.GET("/login", r.UserController.GetUser)
	g.POST("/register", r.UserController.CreateUser)

	g2 := e.Group("")

	conf := middleware.JWTConfig{
		Claims:     &config.JwtCustomClaims{},
		SigningKey: []byte("secret"),
	}
	g2.Use(middleware.JWTWithConfig(conf))
	g2.GET("/auth", r.UserController.CreateUser)
}
