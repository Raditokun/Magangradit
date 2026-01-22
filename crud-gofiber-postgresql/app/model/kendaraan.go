package model

import "time"

type Kendaraan struct {
	ID               int        `json:"id"`
	IdJenisKendaraan int        `json:"id_jenis_kendaraan"`
	Nopol            string     `json:"nopol"`
	NoBpkb           string     `json:"no_bpkb"`
	NoMesin          string     `json:"no_mesin"`
	NoRangka         string     `json:"no_rangka"`
	NamaKendaraan    string     `json:"nama_kendaraan"`
	Warna            string     `json:"warna"`
	Kapasitas        int        `json:"kapasitas"`
	Deskripsi        string     `json:"deskripsi"`
	Status           string     `json:"status"`
	CreatedAt        time.Time  `json:"created_at"`
	CreatedBy        string     `json:"created_by"`
	UpdatedAt        *time.Time `json:"updated_at"`
	UpdatedBy        *string    `json:"updated_by"`
	DeletedAt        *time.Time `json:"deleted_at"`
	DeletedBy        *string    `json:"deleted_by"`
}

type UpdateKendaraanRequest struct {
	IdJenisKendaraan int    `json:"id_jenis_kendaraan" validate:"required"`
	Nopol            string `json:"nopol" validate:"required"`
	NoBpkb           string `json:"no_bpkb"`
	NoMesin          string `json:"no_mesin"`
	NoRangka         string `json:"no_rangka"`
	NamaKendaraan    string `json:"nama_kendaraan" validate:"required"`
	Warna            string `json:"warna"`
	Kapasitas        int    `json:"kapasitas"`
	Deskripsi        string `json:"deskripsi"`
	Status           string `json:"status"`
	UpdatedBy        string `json:"updated_by" validate:"required"`
}

type DeleteKendaraanRequest struct {
	DeletedBy string `json:"deleted_by" validate:"required"`
}
