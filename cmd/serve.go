package cmd

import (
	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"
	routes2 "go-auth/routes"
	"go-auth/server"
)

var serveCmd = &cobra.Command{
	Use: "serve",
	Run: serve,
}

func serve(cmd *cobra.Command, args []string) {
	//TODO: Initiate all clients
	//db = conn.DB()
	//TODO: Initiate Repositories, Services, Controllers

	var echo = echo.New()
	var Routes = routes2.New(echo)
	var Server = server.New(echo)

	Routes.Init()
	Server.Start()
}
