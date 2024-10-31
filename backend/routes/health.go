package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/magdapyndryk/mp-site/backend/controllers"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/api/health", controllers.HealthCheck)
}
