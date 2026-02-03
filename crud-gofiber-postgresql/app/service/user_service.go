package service

import (
	"crud-app/app/model"
	"crud-app/app/repository"
	"crud-app/helper"
	"database/sql"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService() *UserService {
	return &UserService{
		repo: repository.NewUserRepository(),
	}
}

// GetAllUsers godoc
// @Summary Dapatkan semua user
// @Description Mengambil daftar semua user dari database
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {array} model.User
// @Failure 500 {object} helper.Response
// @Router /users [get]
func (s *UserService) GetAllUsers(c *fiber.Ctx) error {
	nip := c.Locals("nip").(string)
	log.Printf("User %s mengakses GET /api/user", nip)

	users, err := s.repo.GetAll()
	if err != nil {
		return helper.Error(c, fiber.StatusInternalServerError, "Gagal mengambil data user")
	}

	return helper.Success(c, users, "Data user berhasil diambil")
}

// HGetUserByID godoc
// @Summary Dapatkan user berdasarkan ID
// @Description Mengambil data user spesifik berdasarkan ID
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} model.User
// @Failure 400 {object} helper.Response
// @Failure 404 {object} helper.Response
// @Router /users/{id} [get]
func (s *UserService) GetUserByID(c *fiber.Ctx) error {
	nip := c.Locals("nip").(string)
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return helper.Error(c, fiber.StatusBadRequest, "ID tidak valid")
	}

	log.Printf("User %s mengakses GET /api/user/%d", nip, id)

	user, err := s.repo.GetByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return helper.Error(c, fiber.StatusNotFound, "User tidak ditemukan")
		}
		return helper.Error(c, fiber.StatusInternalServerError, "Gagal mengambil data user")
	}

	return helper.Success(c, user, "Data user berhasil diambil")
}

// CreateUser godoc
// @Summary Buat user baru
// @Description Membuat user baru di database
// @Tags Users
// @Accept json
// @Produce json
// @Param body body model.CreateUserRequest true "User data"
// @Success 201 {object} model.User
// @Failure 400 {object} helper.Response
// @Failure 500 {object} helper.Response
// @Router /users [post]
func (s *UserService) CreateUser(c *fiber.Ctx) error {
	nip := c.Locals("nip").(string)
	log.Printf("Admin %s menambah user baru", nip)

	var req model.CreateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return helper.Error(c, fiber.StatusBadRequest, "Request body tidak valid")
	}

	if req.Nip == "" || req.Email == "" || req.Role == "" {
		return helper.Error(c, fiber.StatusBadRequest, "Semua field harus diisi")
	}

	user, err := s.repo.Create(req, nip)
	if err != nil {
		return helper.Error(c, fiber.StatusInternalServerError, "Gagal menambah user. Pastikan NIP dan email belum digunakan")
	}

	return c.Status(fiber.StatusCreated).JSON(helper.Response{
		Success: true,
		Message: "User berhasil ditambahkan",
		Data:    user,
	})
}

// UpdateUser godoc
// @Summary Update user
// @Description Memperbarui data user berdasarkan ID
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param body body model.UpdateUserRequest true "User data"
// @Success 200 {object} model.User
// @Failure 400 {object} helper.Response
// @Failure 404 {object} helper.Response
// @Router /users/{id} [put]
func (s *UserService) UpdateUser(c *fiber.Ctx) error {
	nip := c.Locals("nip").(string)
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return helper.Error(c, fiber.StatusBadRequest, "ID tidak valid")
	}

	log.Printf("Admin %s mengupdate User ID %d", nip, id)

	var req model.UpdateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return helper.Error(c, fiber.StatusBadRequest, "Body tidak valid")
	}

	if req.Nip == "" || req.Email == "" || req.Role == "" {
		return helper.Error(c, fiber.StatusBadRequest, "nip, email dan role harus diisi")
	}

	req.UpdatedBy = nip

	user, err := s.repo.Update(id, req)
	if err != nil {
		return helper.Error(c, fiber.StatusInternalServerError, "Gagal mengupdate user")
	}

	if user == nil {
		return helper.Error(c, fiber.StatusNotFound, "User tidak ditemukan")
	}

	return helper.Success(c, user, "User berhasil di update")
}

// DeleteUser godoc
// @Summary Hapus user
// @Description Menghapus user berdasarkan ID
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} helper.Response
// @Failure 400 {object} helper.Response
// @Failure 404 {object} helper.Response
// @Router /users/{id} [delete]
func (s *UserService) DeleteUser(c *fiber.Ctx) error {
	adminNip := c.Locals("nip").(string)
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return helper.Error(c, fiber.StatusBadRequest, "ID tidak valid")
	}

	log.Printf("Admin %s menghapus User ID %d", adminNip, id)

	_, err = s.repo.GetByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return helper.Error(c, fiber.StatusNotFound, "User tidak ditemukan")
		}
		return helper.Error(c, fiber.StatusInternalServerError, "Gagal mengecek user")
	}

	err = s.repo.Delete(id)
	if err != nil {
		log.Printf("Delete error: %v", err)
		return helper.Error(c, fiber.StatusInternalServerError, "Gagal menghapus user")
	}

	return helper.Success(c, nil, "User berhasil dihapus")
}
