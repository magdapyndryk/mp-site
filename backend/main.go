package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/magdapyndryk/mp-site/backend/routes"
)

func main() {
	app := fiber.New()

	routes.RegisterProductRoutes(app)

	app.Listen(":3000")
}
