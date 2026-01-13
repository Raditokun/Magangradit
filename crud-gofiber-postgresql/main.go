package main

import (
	"crud-app/database"
	"crud-app/models"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {

	database.Connect()
	defer database.DB.Close()

	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		},
	})

	// Root route - API info
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

	app.Get("/driver", GetAllDrivers)
	app.Put("/driver/:id", UpdateDriver)
	app.Delete("/driver/:id", SoftDeleteDriver)

	log.Fatal(app.Listen(":3002"))
}

func GetAllDrivers(c *fiber.Ctx) error {
	query := `
		SELECT id, nama, nip, foto, created_at, created_by, updated_at, updated_by 
		FROM driver 
		WHERE deleted_at IS NULL
		ORDER BY id DESC
	`

	rows, err := database.DB.Query(query)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	defer rows.Close()

	var drivers []models.Driver
	for rows.Next() {
		var d models.Driver
		if err := rows.Scan(&d.ID, &d.Nama, &d.NIP, &d.Foto, &d.CreatedAt, &d.CreatedBy, &d.UpdatedAt, &d.UpdatedBy); err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}
		drivers = append(drivers, d)
	}

	return c.JSON(fiber.Map{
		"succecs": true,
		"data":    drivers,
		"message": "Data driver berhasil diambil",
	})
}

func UpdateDriver(c *fiber.Ctx) error {
	id := c.Params("id")

	var req models.UpdateDriverRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Data tidak valid"})
	}

	query := `UPDATE driver SET nama=$1, nip=$2, foto=$3, updated_at=$4, updated_by=$5 WHERE id=$6 AND deleted_at IS NULL`
	result, err := database.DB.Exec(query, req.Nama, req.NIP, req.Foto, time.Now(), req.UpdatedBy, id)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal update driver"})
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{"error": "Driver tidak ditemukan"})
	}

	return c.JSON(fiber.Map{"message": "Driver berhasil diupdate"})
}

func SoftDeleteDriver(c *fiber.Ctx) error {
	id := c.Params("id")

	var req models.DeleteDriverRequest
	if err := c.BodyParser(&req); err != nil {

		req.DeletedBy = ""
	}

	query := `UPDATE driver SET deleted_at=$1, deleted_by=$2 WHERE id=$3 AND deleted_at IS NULL`
	result, err := database.DB.Exec(query, time.Now(), req.DeletedBy, id)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal menghapus driver"})
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{"error": "Driver tidak ditemukan atau sudah dihapus"})
	}

	return c.JSON(fiber.Map{"message": "Driver berhasil dihapus (Soft Delete)"})
}
