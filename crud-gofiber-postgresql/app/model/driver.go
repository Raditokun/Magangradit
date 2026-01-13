package model

import "time"


type Driver struct {
	ID        int        `json:"id"`
	Nama      string     `json:"nama"`
	NIP       *string    `json:"nip"`
	Foto      *string    `json:"foto"`
	CreatedAt time.Time  `json:"created_at"`
	CreatedBy *string    `json:"created_by"`
	UpdatedAt *time.Time `json:"updated_at"`
	UpdatedBy *string    `json:"updated_by"`
	DeletedAt *time.Time `json:"deleted_at"`
	DeletedBy *string    `json:"deleted_by"`
}


type UpdateDriverRequest struct {
	Nama      string `json:"nama"`
	NIP       string `json:"nip"`
	Foto      string `json:"foto"`
	UpdatedBy string `json:"updated_by"`
}


type DeleteDriverRequest struct {
	DeletedBy string `json:"deleted_by"`
}
