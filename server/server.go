package server

import (
	"github.com/labstack/echo/v4"
	"github.com/rakib-09/go-auth/config"
)

type Server struct {
	echo *echo.Echo
}

func (s *Server) Start() {
	e := s.echo
	e.Logger.Fatal(e.Start(":" + config.App().Port))
}

func New(e *echo.Echo) *Server {
	return &Server{
		echo: e,
	}
}
