package routes

import "github.com/labstack/echo/v4"

type Routes struct {
	echo *echo.Echo
}

func New(e *echo.Echo) *Routes {
	return &Routes{
		echo: e,
	}
}

func (r *Routes) Init() {

}
