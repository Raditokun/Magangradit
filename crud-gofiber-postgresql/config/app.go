package config

import (
	"crud-app/middleware"
	"crud-app/route"

	"github.com/gofiber/fiber/v2"
)


func NewApp() *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(500).JSON(fiber.Map{
				"success": false,
				"message": err.Error(),
			})
		},
	})

	//setup
	middleware.SetupMiddleware(app)

	//route
	route.SetupRoutes(app)

	return app
}
