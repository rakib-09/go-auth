package cmd

import (
	"github.com/labstack/echo/v4"
	"github.com/rakib-09/go-auth/controllers"
	"github.com/rakib-09/go-auth/db"
	"github.com/rakib-09/go-auth/repositories"
	r "github.com/rakib-09/go-auth/routes"
	"github.com/rakib-09/go-auth/server"
	"github.com/rakib-09/go-auth/services"
	"github.com/spf13/cobra"
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
	brandRepo := repositories.NewBrandRepository(dbClient)
	// services
	userSvc := services.NewUserService(userRepo)
	jwtSvc := services.NewJwtSvc()
	authSvc := services.NewAuthService(userSvc, jwtSvc)
	companySvc := services.NewCompanyService(companyRepo)
	brandSvc := services.NewBrandService(brandRepo)
	// controllers
	userController := controllers.NewUserController(userSvc)
	authController := controllers.NewAuthController(authSvc)
	companyController := controllers.NewCompanyController(companySvc)
	brandController := controllers.NewBrandController(brandSvc)

	var echo = echo.New()
	var Routes = r.New(echo, userController, authController, companyController, brandController)
	var Server = server.New(echo)

	Routes.Init()
	Server.Start()

	Server.GracefulShutdown(echo)

}
