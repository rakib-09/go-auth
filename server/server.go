package server

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/rakib-09/go-auth/config"
	"os"
	"os/signal"
	"time"
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

func (s *Server) GracefulShutdown(e *echo.Echo) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
