package route

import (
	"crud-app/app/service"
	"crud-app/middleware"

	_ "crud-app/docs" // Import generated docs

	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

func SetupRoutes(app *fiber.App) {
	
	app.Get("/swagger/*", fiberSwagger.WrapHandler)
	driverService := service.NewDriverService()
	kendaraanService := service.NewKendaraanService()
	jenisKendaraanService := service.NewJenisKendaraanService()
	driverKendaraanService := service.NewDriverKendaraanService()
	authService := service.NewAuthService()
	userService := service.NewUserService()
	roleService := service.NewRoleService()

	
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello",
			"endpoints": fiber.Map{
				"POST /api/login":       "Login",
				"GET /api/profile":      "Get profile (protected)",
				"GET /api/user":         "Get all users (protected)",
				"GET /api/role":         "Get all roles (protected)",
				"GET /driver":           "Get all drivers",
				"PUT /driver/:id":       "Update a driver",
				"DELETE /driver/:id":    "Soft delete a driver",
				"GET /kendaraan":        "Get all kendaraan",
				"GET /jenis_kendaraan":  "Get all jenis kendaraan",
				"GET /driver_kendaraan": "Get all driver kendaraan",
			},
		})
	})

	
	api := app.Group("/api")

	
	api.Post("/login", authService.Login)
	api.Get("/hash/:password", authService.HashPassword)

	
	protected := api.Group("", middleware.AuthRequired())
	protected.Get("/profile", authService.GetProfile)

	
	user := protected.Group("/user")
	user.Get("/", userService.GetAllUsers)
	user.Get("/:id", userService.GetUserByID)
	user.Post("/", middleware.AdminOnly(), userService.CreateUser)
	user.Put("/:id", middleware.AdminOnly(), userService.UpdateUser)
	user.Delete("/:id", middleware.AdminOnly(), userService.DeleteUser)

	
	role := protected.Group("/role")
	role.Get("/", roleService.GetAllRoles)
	role.Get("/:id", roleService.GetRoleByID)
	role.Post("/", middleware.AdminOnly(), roleService.CreateRole)
	role.Put("/:id", middleware.AdminOnly(), roleService.UpdateRole)
	role.Delete("/:id", middleware.AdminOnly(), roleService.DeleteRole)

	
	app.Get("/driver", driverService.GetAllDrivers)
	app.Put("/driver/:id", driverService.UpdateDriver)
	app.Delete("/driver/:id", driverService.SoftDeleteDriver)

	
	app.Get("/kendaraan", kendaraanService.GetAllKendaraan)
	app.Put("/kendaraan/:id", kendaraanService.SoftDeleteKendaraan)
	app.Delete("/kendaraan/:id", kendaraanService.SoftDeleteKendaraan)

	
	app.Get("/jenis_kendaraan", jenisKendaraanService.GetAllJenisKendaraan)
	app.Put("/jenis_kendaraan/:id", jenisKendaraanService.UpdateJenisKendaraan)
	app.Delete("/jenis_kendaraan/:id", jenisKendaraanService.SoftDeleteJenisKendaraan)

	app.Get("/driver_kendaraan", driverKendaraanService.GetAllDriverKendaraan)
	app.Put("/driver_kendaraan/:id", driverKendaraanService.UpdateDriverKendaraan)
	app.Delete("/driver_kendaraan/:id", driverKendaraanService.SoftDeleteDriverKendaraan)
}
