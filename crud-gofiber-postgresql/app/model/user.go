package model

import "time"

type User struct {
	ID        int        `json:"id"`
	Role      string     `json:"role"`
	Nip       string     `json:"nip"`
	Email     string     `json:"email"`
	Status    string     `json:"status"`
	CreatedAt time.Time  `json:"created_at"`
	CreatedBy *string    `json:"created_by"`
	UpdatedAt *time.Time `json:"updated_at"`
	UpdatedBy *string    `json:"updated_by"`
}

// UserWithPassword includes password hash for authentication queries
type UserWithPassword struct {
	ID           int        `json:"id"`
	Role         string     `json:"role"`
	Nip          string     `json:"nip"`
	Email        string     `json:"email"`
	PasswordHash string     `json:"-"`
	Status       string     `json:"status"`
	CreatedAt    time.Time  `json:"created_at"`
	CreatedBy    *string    `json:"created_by"`
	UpdatedAt    *time.Time `json:"updated_at"`
	UpdatedBy    *string    `json:"updated_by"`
}

// ToUser converts UserWithPassword to User (without password)
func (u *UserWithPassword) ToUser() User {
	return User{
		ID:        u.ID,
		Role:      u.Role,
		Nip:       u.Nip,
		Email:     u.Email,
		Status:    u.Status,
		CreatedAt: u.CreatedAt,
		CreatedBy: u.CreatedBy,
		UpdatedAt: u.UpdatedAt,
		UpdatedBy: u.UpdatedBy,
	}
}

type CreateUserRequest struct {
	Nip   string `json:"nip"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

type UpdateUserRequest struct {
	Nip       string `json:"nip"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	UpdatedBy string `json:"updated_by"`
}

type DeleteUserRequest struct {
	DeletedBy string `json:"deleted_by"`
}
