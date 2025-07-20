package main

import (
	"fmt"
	"log"
	"os"
	"payslips/bootstrap"
	"payslips/internal/business"
	"payslips/internal/handlers"
	"payslips/internal/middleware"
	"payslips/internal/provider"
	"payslips/internal/repositories"
	"payslips/internal/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// Default config
	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  os.Getenv("APP_NAME"),
		AppName:       os.Getenv("APP_NAME"),
	})

	app.Use(logger.New())

	// Connect to the PostgreSQL database
	db := bootstrap.ConnectDB()
	bootstrap.RunMigrations(db)

	providercfg := bootstrap.Provider()
	provider := provider.NewProvider(providercfg)

	repo := repositories.NewRepository(db)
	business := business.NewBusiness(&repo, provider)
	handler := handlers.NewHandler(business)
	middleware := middleware.NewAuthentication(business)

	routes.Routes(app, handler, middleware)

	port := ":3000"
	if os.Getenv("PORT") != "" {
		port = fmt.Sprintf(":%v", os.Getenv("PORT"))
	}

	log.Println(app.Listen(port))
}
