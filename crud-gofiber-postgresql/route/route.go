package route

import (
	"crud-app/app/service"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	driverService := service.NewDriverService()
	KendaraanService := service.NewKendaraanService()
	JenisKendaraanService := service.NewJenisKendaraanService()
	DriverKendaraanService := service.NewDriverKendaraanService()


	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello",
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

	app.Get("/kendaraan", KendaraanService.GetAllKendaraan)
	app.Put("/kendaraan/:id", KendaraanService.SoftDeleteKendaraan)
	app.Delete("/kendaraan/:id", KendaraanService.SoftDeleteKendaraan)

	app.Get("/jenis_kendaraan", JenisKendaraanService.GetAllJenisKendaraan)
	app.Put("/jenis_kendaraan/:id", JenisKendaraanService.UpdateJenisKendaraan)
	app.Delete("/jenis_kendaraan/:id", JenisKendaraanService.SoftDeleteJenisKendaraan)

	app.Get("/driver_kendaraan", DriverKendaraanService.GetAllDriverKendaraan)
	app.Put("/driver_kendaraan/:id", DriverKendaraanService.UpdateDriverKendaraan)
	app.Delete("/driver_kendaraan/:id", DriverKendaraanService.SoftDeleteDriverKendaraan)


}
	