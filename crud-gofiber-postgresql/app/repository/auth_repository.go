package repository

import (
	"crud-app/app/model"
	"crud-app/database"
)

type AuthRepository interface {
	FindByNipOrEmail(nipOrEmail string) (*model.UserWithPassword, error)
}

type authRepository struct{}

func NewAuthRepository() AuthRepository {
	return &authRepository{}
}

func (r *authRepository) FindByNipOrEmail(nipOrEmail string) (*model.UserWithPassword, error) {
	query := `
		SELECT id, role, nip, email, password_hash, status, created_at, created_by, updated_at, updated_by 
		FROM users
		WHERE nip = $1 OR email = $1
	`

	var u model.UserWithPassword
	err := database.DB.QueryRow(query, nipOrEmail).Scan(
		&u.ID, &u.Role, &u.Nip, &u.Email, &u.PasswordHash, &u.Status, &u.CreatedAt, &u.CreatedBy, &u.UpdatedAt, &u.UpdatedBy,
	)
	if err != nil {
		return nil, err
	}

	return &u, nil
}
