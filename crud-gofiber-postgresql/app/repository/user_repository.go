package repository

import (
	"crud-app/app/model"
	"crud-app/database"
	"time"
)

type UserRepository interface {
	GetAll() ([]model.User, error)
	GetByID(id int) (*model.User, error)
	Create(req model.CreateUserRequest, createdBy string) (*model.User, error)
	Update(id int, req model.UpdateUserRequest) (*model.User, error)
	Delete(id int) error
}

type userRepository struct{}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (r *userRepository) GetAll() ([]model.User, error) {
	query := `
		SELECT id, role, nip, email, status, created_at, created_by, updated_at, updated_by 
		FROM users
		ORDER BY created_at DESC
	`

	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var u model.User
		if err := rows.Scan(&u.ID, &u.Role, &u.Nip, &u.Email, &u.Status, &u.CreatedAt, &u.CreatedBy, &u.UpdatedAt, &u.UpdatedBy); err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}

func (r *userRepository) GetByID(id int) (*model.User, error) {
	query := `
		SELECT id, role, nip, email, status, created_at, created_by, updated_at, updated_by 
		FROM users
		WHERE id = $1
	`

	var u model.User
	err := database.DB.QueryRow(query, id).Scan(
		&u.ID, &u.Role, &u.Nip, &u.Email, &u.Status, &u.CreatedAt, &u.CreatedBy, &u.UpdatedAt, &u.UpdatedBy,
	)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (r *userRepository) Create(req model.CreateUserRequest, createdBy string) (*model.User, error) {
	now := time.Now()
	query := `
		INSERT INTO users (nip, role, email, created_at, created_by, updated_at, updated_by)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id
	`

	var id int
	err := database.DB.QueryRow(query, req.Nip, req.Role, req.Email, now, createdBy, now, createdBy).Scan(&id)
	if err != nil {
		return nil, err
	}

	return r.GetByID(id)
}

func (r *userRepository) Update(id int, req model.UpdateUserRequest) (*model.User, error) {
	query := `
		UPDATE users 
		SET nip = $1, role = $2, email = $3, updated_at = $4, updated_by = $5
		WHERE id = $6
	`

	result, err := database.DB.Exec(query, req.Nip, req.Role, req.Email, time.Now(), req.UpdatedBy, id)
	if err != nil {
		return nil, err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return nil, nil
	}

	return r.GetByID(id)
}

func (r *userRepository) Delete(id int) error {
	// Check if user exists first
	var exists bool
	err := database.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE id = $1)", id).Scan(&exists)
	if err != nil {
		return err
	}
	if !exists {
		return nil
	}

	_, err = database.DB.Exec("DELETE FROM users WHERE id = $1", id)
	return err
}
