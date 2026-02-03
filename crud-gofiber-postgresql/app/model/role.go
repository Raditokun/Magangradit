package model

import "time"

type Role struct {
	ID        int        `json:"id"`
	Name      string     `json:"name"`
	CreatedAt time.Time  `json:"created_at"`
	CreatedBy *string    `json:"created_by"`
	UpdatedAt *time.Time `json:"updated_at"`
	UpdatedBy *string    `json:"updated_by"`
}

type CreateRoleRequest struct {
	Name      string `json:"name"`
	CreatedBy string `json:"created_by"`
}

type UpdateRoleRequest struct {
	Name      string `json:"name"`
	UpdatedBy string `json:"updated_by"`
}

type DeleteRoleRequest struct {
	DeletedBy string `json:"deleted_by"`
}
