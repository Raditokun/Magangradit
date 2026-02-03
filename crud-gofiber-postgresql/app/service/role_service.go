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

type RoleService struct {
	repo repository.RoleRepository
}

func NewRoleService() *RoleService {
	return &RoleService{
		repo: repository.NewRoleRepository(),
	}
}

// GetAllRoles returns all roles
func (s *RoleService) GetAllRoles(c *fiber.Ctx) error {
	nip := c.Locals("nip").(string)
	log.Printf("User %s mengakses GET /api/role", nip)

	roles, err := s.repo.GetAll()
	if err != nil {
		return helper.Error(c, fiber.StatusInternalServerError, "Gagal mengambil data role")
	}

	return helper.Success(c, roles, "Data role berhasil diambil")
}

// GetRoleByID returns a single role by ID
func (s *RoleService) GetRoleByID(c *fiber.Ctx) error {
	nip := c.Locals("nip").(string)
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return helper.Error(c, fiber.StatusBadRequest, "ID tidak valid")
	}

	log.Printf("User %s mengakses GET /api/role/%d", nip, id)

	role, err := s.repo.GetByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return helper.Error(c, fiber.StatusNotFound, "Role tidak ditemukan")
		}
		return helper.Error(c, fiber.StatusInternalServerError, "Gagal mengambil data role")
	}

	return helper.Success(c, role, "Data role berhasil diambil")
}

// CreateRole creates a new role (admin only)
func (s *RoleService) CreateRole(c *fiber.Ctx) error {
	nip := c.Locals("nip").(string)
	log.Printf("Admin %s menambah role baru", nip)

	var req model.CreateRoleRequest
	if err := c.BodyParser(&req); err != nil {
		return helper.Error(c, fiber.StatusBadRequest, "Request body tidak valid")
	}

	if req.Name == "" {
		return helper.Error(c, fiber.StatusBadRequest, "Nama role harus diisi")
	}

	req.CreatedBy = nip

	role, err := s.repo.Create(req)
	if err != nil {
		return helper.Error(c, fiber.StatusInternalServerError, "Gagal menambah role")
	}

	return c.Status(fiber.StatusCreated).JSON(helper.Response{
		Success: true,
		Message: "Role berhasil ditambahkan",
		Data:    role,
	})
}

// UpdateRole updates an existing role (admin only)
func (s *RoleService) UpdateRole(c *fiber.Ctx) error {
	nip := c.Locals("nip").(string)
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return helper.Error(c, fiber.StatusBadRequest, "ID tidak valid")
	}

	log.Printf("Admin %s mengupdate Role ID %d", nip, id)

	var req model.UpdateRoleRequest
	if err := c.BodyParser(&req); err != nil {
		return helper.Error(c, fiber.StatusBadRequest, "Body tidak valid")
	}

	if req.Name == "" {
		return helper.Error(c, fiber.StatusBadRequest, "Nama role harus diisi")
	}

	req.UpdatedBy = nip

	role, err := s.repo.Update(id, req)
	if err != nil {
		return helper.Error(c, fiber.StatusInternalServerError, "Gagal mengupdate role")
	}

	if role == nil {
		return helper.Error(c, fiber.StatusNotFound, "Role tidak ditemukan")
	}

	return helper.Success(c, role, "Role berhasil di update")
}

// DeleteRole soft deletes a role (admin only)
func (s *RoleService) DeleteRole(c *fiber.Ctx) error {
	adminNip := c.Locals("nip").(string)
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return helper.Error(c, fiber.StatusBadRequest, "ID tidak valid")
	}

	log.Printf("Admin %s menghapus Role ID %d", adminNip, id)

	// Check if role exists
	_, err = s.repo.GetByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return helper.Error(c, fiber.StatusNotFound, "Role tidak ditemukan")
		}
		return helper.Error(c, fiber.StatusInternalServerError, "Gagal mengecek role")
	}

	err = s.repo.Delete(id)
	if err != nil {
		log.Printf("Delete error: %v", err)
		return helper.Error(c, fiber.StatusInternalServerError, "Gagal menghapus role")
	}

	return helper.Success(c, nil, "Role berhasil dihapus")
}
