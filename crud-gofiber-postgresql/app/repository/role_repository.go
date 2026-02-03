package repository

import (
	"crud-app/app/model"
	"crud-app/database"
	"time"
)

type RoleRepository interface {
	GetAll() ([]model.Role, error)
	GetByID(id int) (*model.Role, error)
	Create(req model.CreateRoleRequest) (*model.Role, error)
	Update(id int, req model.UpdateRoleRequest) (*model.Role, error)
	Delete(id int) error
}

type roleRepository struct{}

func NewRoleRepository() RoleRepository {
	return &roleRepository{}
}

func (r *roleRepository) GetAll() ([]model.Role, error) {
	query := `
		SELECT id, name, created_at, created_by, updated_at, updated_by 
		FROM roles
		WHERE deleted_at IS NULL
		ORDER BY id ASC
	`

	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var roles []model.Role
	for rows.Next() {
		var role model.Role
		if err := rows.Scan(&role.ID, &role.Name, &role.CreatedAt, &role.CreatedBy, &role.UpdatedAt, &role.UpdatedBy); err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}

	return roles, nil
}

func (r *roleRepository) GetByID(id int) (*model.Role, error) {
	query := `
		SELECT id, name, created_at, created_by, updated_at, updated_by 
		FROM roles
		WHERE id = $1 AND deleted_at IS NULL
	`

	var role model.Role
	err := database.DB.QueryRow(query, id).Scan(
		&role.ID, &role.Name, &role.CreatedAt, &role.CreatedBy, &role.UpdatedAt, &role.UpdatedBy,
	)
	if err != nil {
		return nil, err
	}

	return &role, nil
}

func (r *roleRepository) Create(req model.CreateRoleRequest) (*model.Role, error) {
	now := time.Now()
	query := `
		INSERT INTO roles (name, created_at, created_by, updated_at, updated_by)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`

	var id int
	err := database.DB.QueryRow(query, req.Name, now, req.CreatedBy, now, req.CreatedBy).Scan(&id)
	if err != nil {
		return nil, err
	}

	return r.GetByID(id)
}

func (r *roleRepository) Update(id int, req model.UpdateRoleRequest) (*model.Role, error) {
	query := `
		UPDATE roles 
		SET name = $1, updated_at = $2, updated_by = $3
		WHERE id = $4 AND deleted_at IS NULL
	`

	result, err := database.DB.Exec(query, req.Name, time.Now(), req.UpdatedBy, id)
	if err != nil {
		return nil, err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return nil, nil
	}

	return r.GetByID(id)
}

func (r *roleRepository) Delete(id int) error {
	query := `
		UPDATE roles 
		SET deleted_at = $1 
		WHERE id = $2 AND deleted_at IS NULL
	`

	_, err := database.DB.Exec(query, time.Now(), id)
	return err
}
