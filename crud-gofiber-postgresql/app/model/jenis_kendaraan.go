package model

import "time"

type JenisKendaraan struct{
	ID int `json:"id"`
	Nama string `json:"nama"`
	CreatedAt time.Time  `json:"created_at"`
	CreatedBy *string    `json:"created_by"`
	UpdatedAt *time.Time `json:"updated_at"`
	UpdatedBy *string    `json:"updated_by"`
	DeletedAt *time.Time `json:"deleted_at"`
	DeletedBy *string    `json:"deleted_by"`


}


type UpdateJenisKendaraanRequest struct {
	Nama string `json:"nama"`
	UpdatedBy string `json:"updated_by"`
}

type DeleteJenisKendaraanRequest struct {
	DeletedBy string `json:"deleted_by"`

}