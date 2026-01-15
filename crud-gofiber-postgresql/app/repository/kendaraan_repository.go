package repository

import (
	"crud-app/app/model"
	"crud-app/database"
	"time"
)

type KendaraanRepository interface {
	GetAll() ([]model.Kendaraan, error)
	Update(id string, req model.UpdateKendaraanRequest) (int64, error)
	SoftDelete(id string, req model.DeleteKendaraanRequest) (int64, error)
}

type kendaraanRepository struct{}

func NewKendaraanRepository() KendaraanRepository {
	return &kendaraanRepository{}
}

func (r *kendaraanRepository) GetAll() ([]model.Kendaraan, error) {
	query := `
	SELECT id, id_jenis_kendaraan, nopol, no_bpkb, no_mesin, nama_kendaraan, warna, kapasitas, status,created_at,created_by, updated_at, updated_by
	FROM kendaraan
	WHERE deleted_at IS NULL
	ORDER BY id DESC`

	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var kendaraan []model.Kendaraan
	for rows.Next() {
		var k model.Kendaraan
		if err := rows.Scan(&k.ID, &k.ID_jenis_kendaraan, &k.Nopol, &k.No_bpkb, &k.No_mesin, &k.Nama_Kendaraan, &k.Warna, &k.Kapasitas, &k.Deskripsi, &k.Status, &k.CreatedAt, &k.CreatedBy, &k.UpdatedAt, &k.UpdatedBy); err != nil {
			return nil, err
		}
		kendaraan = append(kendaraan, k)
	}

	return kendaraan, nil // ke model kendaraan
}

func (r *kendaraanRepository) Update(id string, req model.UpdateKendaraanRequest) (int64, error) {
	query := `UPDATE kendaraan SET id_jenis_kendaraan=$1, nopol=$2, no_bpkb=$3, no_mesin=$4, nama_kendaraan=$5, warna=$6, kapasitas=$7, deskripsi=$8, status=$9, updated_at=$10, updated_by=$11 WHERE id=$12 AND deleted_at IS NULL`
	result, err := database.DB.Exec(query, req.ID_jenis_kendaraan, req.Nopol, req.No_bpkb, req.No_mesin, req.Nama_Kendaraan, req.Warna, req.Kapasitas, req.Deskripsi, req.Status, time.Now(), req.UpdatedBy, id)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (r *kendaraanRepository) SoftDelete(id string, req model.DeleteKendaraanRequest) (int64, error) {
	query := `UPDATE kendaraan SET deleted_at=$1, deleted_by=$2 WHERE id=$3 AND deleted_at IS NULL`
	result, err := database.DB.Exec(query, time.Now(), req.DeletedBy, id)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}
