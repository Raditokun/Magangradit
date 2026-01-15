package service
import (
	"crud-app/app/model"
	"crud-app/app/repository"
	"crud-app/helper"

	"github.com/gofiber/fiber/v2"
)

type JenisKendaraanService struct {
	repo repository.JenisKendaraanRepository
}

func NewJenisKendaraanService() *JenisKendaraanService {
	return &JenisKendaraanService{
		repo: repository.NewJenisKendaraanRepository(),
	}
}

func (s *JenisKendaraanService) GetAllJenisKendaraan(c *fiber.Ctx) error {
	jenis_kendaraan, err := s.repo.GetAll()
	if err != nil {
		return helper.Error(c, fiber.StatusInternalServerError, err.Error())
	}
	return helper.Success(c, jenis_kendaraan, "Data jenis kendaraan berhasil diambil")
}

func (s *JenisKendaraanService) UpdateJenisKendaraan(c *fiber.Ctx) error {
	id := c.Params("id")

	var req model.UpdateJenisKendaraanRequest
	if err := c.BodyParser(&req); err != nil {
		return helper.Error(c, fiber.StatusBadRequest, "Data tidak valid")
	}

	rowsAffected, err := s.repo.Update(id, req)
	if err != nil {
		return helper.Error(c, fiber.StatusInternalServerError, "Gagal update jenis kendaraan")
	}

	if rowsAffected == 0 {
		return helper.Error(c, fiber.StatusNotFound, "Jenis kendaraan tidak ditemukan")
	}

	return helper.Success(c, nil, "Jenis kendaraan berhasil diupdate")
}

func (s *JenisKendaraanService) SoftDeleteJenisKendaraan(c *fiber.Ctx) error {
	id := c.Params("id")

	var req model.DeleteJenisKendaraanRequest
	if err := c.BodyParser(&req); err != nil {
		req.DeletedBy = ""
	}

	rowsAffected, err := s.repo.SoftDelete(id, req)
	if err != nil {
		return helper.Error(c, fiber.StatusInternalServerError, "Gagal menghapus jenis kendaraan")
	}

	if rowsAffected == 0 {
		return helper.Error(c, fiber.StatusNotFound, "Jenis kendaraan tidak ditemukan")
	}
	return helper.Success(c, nil, "Jenis kendaraan berhasil dihapus")

}