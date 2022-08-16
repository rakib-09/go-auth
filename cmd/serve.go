package cmd

import (
	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"
	"go-auth/conn"
	"go-auth/controllers"
	"go-auth/repositories"
	routes2 "go-auth/routes"
	"go-auth/server"
	"go-auth/services"
)

var serveCmd = &cobra.Command{
	Use: "serve",
	Run: serve,
}

func serve(cmd *cobra.Command, args []string) {
	//TODO: Initiate all clients
	dbClient := conn.DB()
	//TODO: Initiate repositories, Services, Controllers
	userRepo := repositories.NewUserRepository(dbClient)
	authSvc := services.NewAuthService(userRepo)

	userController := controllers.NewUserController()
	authController := controllers.NewAuthController(authSvc)

	var echo = echo.New()
	var Routes = routes2.New(echo, userController, authController)
	var Server = server.New(echo)

	Routes.Init()
	Server.Start()
}
