package routes

import "github.com/gofiber/fiber/v2"

func RegisterProductRoutes(app *fiber.App) {
	app.Get("/api/products", getProducts)
}

func getProducts(c *fiber.Ctx) error {
	products := []fiber.Map{
		{"id": 1, "name": "Product 1", "price": 100},
		{"id": 2, "name": "Product 2", "price": 200},
	}
	return c.JSON(products)
}
