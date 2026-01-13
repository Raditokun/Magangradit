package route

import (
	"crud-app/app/service"

	"github.com/gofiber/fiber/v2"
)


func SetupRoutes(app *fiber.App) {
	
	driverService := service.NewDriverService()


	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Welcome to Driver CRUD API",
			"endpoints": fiber.Map{
				"GET /driver":        "Get all drivers",
				"PUT /driver/:id":    "Update a driver",
				"DELETE /driver/:id": "Soft delete a driver",
			},
		})
	})

	app.Get("/driver", driverService.GetAllDrivers)
	app.Put("/driver/:id", driverService.UpdateDriver)
	app.Delete("/driver/:id", driverService.SoftDeleteDriver)
}
