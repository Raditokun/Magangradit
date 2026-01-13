package service

import (
	"crud-app/app/model"
	"crud-app/app/repository"
	"crud-app/helper"

	"github.com/gofiber/fiber/v2"
)


type DriverService struct {
	repo repository.DriverRepository
}


func NewDriverService() *DriverService {
	return &DriverService{
		repo: repository.NewDriverRepository(),
	}
}


func (s *DriverService) GetAllDrivers(c *fiber.Ctx) error {
	drivers, err := s.repo.GetAll()
	if err != nil {
		return helper.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return helper.Success(c, drivers, "Data driver berhasil diambil")
}


func (s *DriverService) UpdateDriver(c *fiber.Ctx) error {
	id := c.Params("id")

	var req model.UpdateDriverRequest
	if err := c.BodyParser(&req); err != nil {
		return helper.Error(c, fiber.StatusBadRequest, "Data tidak valid")
	}

	rowsAffected, err := s.repo.Update(id, req)
	if err != nil {
		return helper.Error(c, fiber.StatusInternalServerError, "Gagal update driver")
	}

	if rowsAffected == 0 {
		return helper.Error(c, fiber.StatusNotFound, "Driver tidak ditemukan")
	}

	return helper.Success(c, nil, "Driver berhasil diupdate")
}


func (s *DriverService) SoftDeleteDriver(c *fiber.Ctx) error {
	id := c.Params("id")

	var req model.DeleteDriverRequest
	if err := c.BodyParser(&req); err != nil {
		req.DeletedBy = ""
	}

	rowsAffected, err := s.repo.SoftDelete(id, req)
	if err != nil {
		return helper.Error(c, fiber.StatusInternalServerError, "Gagal menghapus driver")
	}

	if rowsAffected == 0 {
		return helper.Error(c, fiber.StatusNotFound, "Driver tidak ditemukan atau sudah dihapus")
	}

	return helper.Success(c, nil, "Driver berhasil dihapus (Soft Delete)")
}
