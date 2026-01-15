package service

import (
	"crud-app/app/model"
	"crud-app/app/repository"
	"crud-app/helper"

	"github.com/gofiber/fiber/v2"
)

type DriverKendaraanService struct {
	repo repository.DriverKendaraanRepository
}

func NewDriverKendaraanService() *DriverKendaraanService {
	return &DriverKendaraanService{
		repo: repository.NewDriverKendaraanRepository(),
	}
}

func (s *DriverKendaraanService) GetAllDriverKendaraan(c *fiber.Ctx) error {
	driver_kendaraan, err := s.repo.GetAll()
	if err != nil {
		return helper.Error(c, fiber.StatusInternalServerError, err.Error())
	}
	return helper.Success(c, driver_kendaraan, "Data driver_kendaraan berhasil diambil")

}

func (s *DriverKendaraanService) UpdateDriverKendaraan(c *fiber.Ctx) error {
	id := c.Params("id")

	var req model.UpdateDriverKendaraanRequest
	if err := c.BodyParser(&req); err != nil {
		return helper.Error(c, fiber.StatusBadRequest, "Data tidak valid")
	}
	rowsAffected, err := s.repo.Update(id, req)
	if err != nil {
		return helper.Error(c, fiber.StatusInternalServerError, "Gagal update driver_kendaraan")
	}
	if rowsAffected == 0 {
		return helper.Error(c, fiber.StatusNotFound, "Driver_Kendaraan Not found")

	}
	return helper.Success(c, nil, "Driver_Kendaraan berhasil diupdate")

}

func (s *DriverKendaraanService) SoftDeleteDriverKendaraan(c *fiber.Ctx) error {
	id := c.Params("id")

	var req model.DeleteDriverKendaraanRequest
	if err := c.BodyParser(&req); err != nil {
		return helper.Error(c, fiber.StatusBadRequest, "Data tidak valid")
	}
	rowsAffected, err := s.repo.SoftDelete(id, req)
	if err != nil {
		return helper.Error(c, fiber.StatusInternalServerError, "Gagal delete driver_kendaraan")
	}
	if rowsAffected == 0 {
		return helper.Error(c, fiber.StatusNotFound, "Driver_Kendaraan Not found")
	}
	return helper.Success(c, nil, "Driver_Kendaraan berhasil dihapus")
}
