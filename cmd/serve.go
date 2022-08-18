package cmd

import (
	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"
	"go-auth/controllers"
	"go-auth/db"
	"go-auth/repositories"
	r "go-auth/routes"
	"go-auth/server"
	"go-auth/services"
)

var serveCmd = &cobra.Command{
	Use: "serve",
	Run: serve,
}

func serve(cmd *cobra.Command, args []string) {
	// clients
	dbClient := db.Client()
	// repos
	userRepo := repositories.NewUserRepository(dbClient)
	companyRepo := repositories.NewCompanyRepository(dbClient)
	// services
	userSvc := services.NewUserService(userRepo)
	jwtSvc := services.NewJwtSvc()
	authSvc := services.NewAuthService(userSvc, jwtSvc)
	companySvc := services.NewCompanyService(companyRepo)
	// controllers
	userController := controllers.NewUserController(userSvc)
	authController := controllers.NewAuthController(authSvc)
	companyController := controllers.NewCompanyController(companySvc)

	var echo = echo.New()
	var Routes = r.New(echo, userController, authController, companyController)
	var Server = server.New(echo)

	Routes.Init()
	Server.Start()
}
