package service

import (
    "crud-app/app/model"
	"crud-app/app/repository"
	"crud-app/helper"

	"github.com/gofiber/fiber/v2"
)

type KendaraanService struct {
	repo repository.KendaraanRepository

}

//gives KendaraaanServiece 
func NewKendaraanService() *KendaraanService {
	return &KendaraanService{
		repo: repository.NewKendaraanRepository(),
	}

}

func (s *KendaraanService) GetAllKendaraan(c *fiber.Ctx) error{
	id := c.Params("id")

	var req model.UpdateKendaraanRequest
	if err := c.BodyParser(&req); err != nil {
		return helper.Error(c, fiber.StatusBadRequest, "Data tidak valid")
}

rowsAffected, err := s.repo.Update(id, req)	
if err != nil{
	return helper.Error(c, fiber.StatusInternalServerError, "Gagal update kendaraan")
}

if rowsAffected == 0 {
	return helper.Error(c, fiber.StatusNotFound, "Kendaraan tidak ditemukan")
}

return helper.Success(c, nil, "Kendaraan berhasil diupdate")
}

func (s *KendaraanService) SoftDeleteKendaraan(c *fiber.Ctx) error{
	id := c.Params("id")
	var req model.DeleteKendaraanRequest
	if err := c.BodyParser(&req); err != nil {
		req.DeletedBy = ""	
	}

	rowsAffected, err := s.repo.SoftDelete(id, req)
	if err != nil{
		return helper.Error(c, fiber.StatusInternalServerError, "Gagal menghapus kendaraan")


	}

	if rowsAffected == 0 {
		return helper.Error(c, fiber.StatusNotFound, "Kendaraan tidak ditemukan")
	}
	return helper.Success(c, nil, "Kendaraan berhasil dihapus")
}