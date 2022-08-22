package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rakib-09/go-auth/config"
	c "github.com/rakib-09/go-auth/controllers"
	m "github.com/rakib-09/go-auth/middlewares"
)

type Routes struct {
	echo              *echo.Echo
	UserController    *c.UserController
	AuthController    *c.AuthController
	CompanyController *c.CompanyController
	BrandController   *c.BrandController
}

func New(e *echo.Echo, uc *c.UserController, ac *c.AuthController, cc *c.CompanyController, bc *c.BrandController) *Routes {
	return &Routes{
		echo:              e,
		UserController:    uc,
		AuthController:    ac,
		CompanyController: cc,
		BrandController:   bc,
	}
}

func (r *Routes) Init() {
	e := r.echo
	m.Init(r.echo)
	g := e.Group("/api/v1")

	// public
	g.POST("/login", r.AuthController.Login)
	g.POST("/register", r.UserController.CreateUser)

	g2 := e.Group("/api/v1")

	conf := middleware.JWTConfig{
		Claims:     &config.JwtCustomClaims{},
		SigningKey: []byte("accesstokensecret"),
	}
	g2.Use(middleware.JWTWithConfig(conf))
	g2.GET("/users", r.UserController.GetUser)
	g2.GET("/companies", r.CompanyController.Show)
	g2.POST("/companies", r.CompanyController.Create)
	g2.PUT("/companies/:id", r.CompanyController.Update)
	g2.POST("/brands", r.BrandController.Create)
}
