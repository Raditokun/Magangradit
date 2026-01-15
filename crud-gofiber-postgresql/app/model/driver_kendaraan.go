package model

import "time"


type DriverKendaraan struct{
	ID int `json:"id"`
	ID_kendaraan int `json:"id_kendaraan"`
	ID_driver int `json:"id_driver"`
	CreatedAt time.Time  `json:"created_at"`
	CreatedBy *string    `json:"created_by"`
	UpdatedAt *time.Time `json:"updated_at"`
	UpdatedBy *string    `json:"updated_by"`
	DeletedAt *time.Time `json:"deleted_at"`
	DeletedBy *string    `json:"deleted_by"`
}

type UpdateDriverKendaraanRequest struct {
	ID_kendaraan int `json:"id_kendaraan"`
	ID_driver int `json:"id_driver"`
	UpdatedBy string    `json:"updated_by"`

}

type DeleteDriverKendaraanRequest struct {
	DeletedBy string `json:"deleted_by"`
}