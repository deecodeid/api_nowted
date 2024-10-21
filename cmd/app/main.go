package main

import (
	"fmt"
	"log"

	"github.com/deecodeid/api_nowted/config"
	"github.com/deecodeid/api_nowted/domain/usecases"
	"github.com/deecodeid/api_nowted/helper"
	"github.com/deecodeid/api_nowted/infrastructure/api/middleware"
	"github.com/deecodeid/api_nowted/infrastructure/api/routes"
	"github.com/deecodeid/api_nowted/infrastructure/database"
	"github.com/deecodeid/api_nowted/repository"
	"github.com/deecodeid/api_nowted/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	config.Load()
	database.Load()

	app.Use(logger.New())
	app.Use(middleware.ApiKeyMiddleware)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello, World!",
		})
	})

	// repository
	userRepo := repository.NewUserRepository(database.DB)
	tokenRepo := repository.NewTokenVerificationRepository(database.DB)

	// service
	authService := service.NewAuthService(database.DB, userRepo, tokenRepo)

	// usecase
	authUseCase := usecases.NewAuthUseCase(authService)

	// handler & route
	api := app.Group("/api")

	authHandler := routes.NewAuthRoute(authUseCase)
	routes.SetupAuthRoute(api, authHandler)

	app.Use(func(c *fiber.Ctx) error {
		return helper.HandleResponse(c, 404, "Route is not found", nil)
	})

	APP_LISTEN := fmt.Sprintf(":%s", config.ENV.AppPort)
	log.Fatal(app.Listen(APP_LISTEN))
}
