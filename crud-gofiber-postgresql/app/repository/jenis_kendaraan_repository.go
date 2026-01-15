package repository

import (
	"crud-app/app/model"
	"crud-app/database"
	"time"
)

type JenisKendaraanRepository interface {
	GetAll() ([]model.JenisKendaraan, error)
	Update(id string, req model.UpdateJenisKendaraanRequest) (int64, error)
	SoftDelete(id string, req model.DeleteJenisKendaraanRequest) (int64, error)
}

type jeniskendaraanRepository struct{}

func NewJenisKendaraanRepository() JenisKendaraanRepository {
	return &jeniskendaraanRepository{}
}

func (r *jeniskendaraanRepository) GetAll() ([]model.JenisKendaraan, error) {
	query := `
	SELECT id, nama, created_at,created_by, updated_at, updated_by
	FROM jenis_kendaraan
	WHERE deleted_at IS NULL
	ORDER BY id DESC`

	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var jenis_kendaraan []model.JenisKendaraan
	for rows.Next() {
		var jk model.JenisKendaraan
		if err := rows.Scan(&jk.ID, &jk.Nama, &jk.CreatedAt, &jk.CreatedBy, &jk.UpdatedAt, &jk.UpdatedBy); err != nil {
			return nil, err
		}
		jenis_kendaraan = append(jenis_kendaraan, jk)
	}

	return jenis_kendaraan, nil

}

func (r *jeniskendaraanRepository) Update(id string, req model.UpdateJenisKendaraanRequest) (int64, error) {
	query := `UPDATE jenis_kendaraan SET nama=$1, updated_at=$2, updated_by=$3 WHERE id=$4 AND deleted_at IS NULL`
	result, err := database.DB.Exec(query, req.Nama, time.Now(), req.UpdatedBy, id)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()

}

func (r *jeniskendaraanRepository) SoftDelete(id string, req model.DeleteJenisKendaraanRequest) (int64, error) {
	query := `UPDATE jenis_kendaraan SET deleted_at=$1, deleted_by=$2 WHERE id=$3 AND deleted_at IS NULL`
	result, err := database.DB.Exec(query, time.Now(), req.DeletedBy, id)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}
