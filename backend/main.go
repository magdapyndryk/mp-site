package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/magdapyndryk/mp-site/backend/routes"
)

func main() {
	app := fiber.New()

	routes.SetupRoutes(app)

	app.Listen(":3000")
}
