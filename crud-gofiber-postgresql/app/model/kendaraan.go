package model

import "time"


type Kendaraan struct{
	ID 	  int        `json:"id"`
	ID_jenis_kendaraan int        `json:"id_jenis_kendaraan"`
	Nopol string 	  `json:"nopol"`
	No_bpkb string 	  `json:"no_bpkb"`
	No_mesin string 	  `json:"no_mesin"`
	Nama_Kendaraan string 	  `json:"nama_kendaraan"`
	Warna string 	  `json:"warna"`
	Kapasitas int 	  `json:"kapasitas"`
    Deskripsi string    `json:"deskripsi"`
	Status string    `json:"status"`
	CreatedAt time.Time  `json:"created_at"`
	CreatedBy *string    `json:"created_by"`
	UpdatedAt *time.Time `json:"updated_at"`
	UpdatedBy *string    `json:"updated_by"`
	DeletedAt *time.Time `json:"deleted_at"`
	DeletedBy *string    `json:"deleted_by"`


}

type UpdateKendaraanRequest struct {
	ID_jenis_kendaraan int        `json:"id_jenis_kendaraan"`
	Nopol string 	  `json:"nopol"`
	No_bpkb string 	  `json:"no_bpkb"`
	No_mesin string 	  `json:"no_mesin"`
	Nama_Kendaraan string 	  `json:"nama_kendaraan"`
	Warna string 	  `json:"warna"`
	Kapasitas int 	  `json:"kapasitas"`
    Deskripsi string    `json:"deskripsi"`
	Status string    `json:"status"`
	UpdatedBy string `json:"updated_by"`
}

type DeleteKendaraanRequest struct {
	DeletedBy string `json:"deleted_by"`
}


