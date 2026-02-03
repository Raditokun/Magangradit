package service

import (
	"crud-app/app/model"
	"crud-app/app/repository"
	"crud-app/helper"
	"crud-app/utils"
	"database/sql"
	"log"

	"github.com/gofiber/fiber/v2"
)

type AuthService struct {
	repo repository.AuthRepository
}

func NewAuthService() *AuthService {
	return &AuthService{
		repo: repository.NewAuthRepository(),
	}
}

// Login authenticates a user and returns a JWT token
func (s *AuthService) Login(c *fiber.Ctx) error {
	log.Printf("DEBUG: Raw body: %q", string(c.Body()))
	log.Printf("DEBUG: Content-Type header: %q", c.Get("Content-Type"))

	var req model.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		log.Printf("DEBUG: BodyParser error: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "request body gagal",
			"details": err.Error(),
		})
	}

	log.Printf("DEBUG: Parsed req: %+v", req)

	if req.Nip == "" || req.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Nip/Password tak ada",
		})
	}

	userWithPassword, err := s.repo.FindByNipOrEmail(req.Nip)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Nip atau Password Salah",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error database",
		})
	}

	log.Printf("DEBUG: Found user with NIP: %s, Hash from DB: %s", userWithPassword.Nip, userWithPassword.PasswordHash)
	log.Printf("DEBUG: Password from request: %s", req.Password)

	if !utils.CheckPasswordHash(req.Password, userWithPassword.PasswordHash) {
		log.Printf("DEBUG: Password check FAILED")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Username atau password salah",
		})
	}

	// Convert to User (without password)
	user := userWithPassword.ToUser()

	token, err := utils.GenerateToken(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Gagal generate token",
		})
	}

	response := model.LoginResponse{
		User:  user,
		Token: token,
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Login Berhasil",
		"data":    response,
	})
}

// GetProfile returns the authenticated user's profile
func (s *AuthService) GetProfile(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(int)
	nip := c.Locals("nip").(string)
	role := c.Locals("role").(string)

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Profile berhasil diambil",
		"data": fiber.Map{
			"user_id": userID,
			"nip":     nip,
			"role":    role,
		},
	})
}

// HashPassword is a development utility to generate password hashes
func (s *AuthService) HashPassword(c *fiber.Ctx) error {
	password := c.Params("password")
	hash, err := utils.HashPassword(password)
	if err != nil {
		return helper.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(fiber.Map{
		"password": password,
		"hash":     hash,
	})
}
