package routes

import (
	"github.com/labstack/echo/v4"
	m "go-auth/middlewares"
)

type Routes struct {
	echo *echo.Echo
}

func New(e *echo.Echo) *Routes {
	return &Routes{
		echo: e,
	}
}

func (r *Routes) Init() {
	m.Init(r.echo)
}
