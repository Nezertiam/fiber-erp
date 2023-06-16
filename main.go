package main

import (
	"github.com/joho/godotenv"
	"github.com/nezertiam/fiber-erp/infrastructure"
	"github.com/nezertiam/fiber-erp/internals/core/services"
	"github.com/nezertiam/fiber-erp/internals/handlers"
	"github.com/nezertiam/fiber-erp/internals/middlewares"
	"github.com/nezertiam/fiber-erp/internals/repositories"
	"github.com/nezertiam/fiber-erp/internals/server"
)

// @title Fiber ERP API
// @version 0.0.0
// @description API for an ERP application.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8000
// @BasePath /
// @schemes http
func main() {
	//env
	godotenv.Load()
	//database
	infrastructure.InitDB()
	infrastructure.MigrateAll()
	db := infrastructure.GetDB()
	//repositories
	userRepository := repositories.NewUserRepository(db)
	//services
	userService := services.NewUserService(userRepository)
	//handlers
	userHandlers := handlers.NewUserHandlers(userService)
	//middlewares
	middlewares := middlewares.NewMiddlewares()
	//server
	httpServer := server.NewServer(userHandlers, middlewares)
	httpServer.Initialize()
}
