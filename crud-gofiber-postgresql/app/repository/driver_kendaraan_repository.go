package repository

import (
	"crud-app/app/model"
	"crud-app/database"
	"time"
)

type DriverKendaraanRepository interface {
	GetAll() ([]model.DriverKendaraan, error)
	Update(id string, req model.UpdateDriverKendaraanRequest) (int64, error)
	SoftDelete(id string, req model.DeleteDriverKendaraanRequest) (int64, error)
}

type driver_kendaraanRepository struct{}

func NewDriverKendaraanRepository() DriverKendaraanRepository {
	return &driver_kendaraanRepository{}
}

func (r *driver_kendaraanRepository) GetAll() ([]model.DriverKendaraan, error) {
	query := `
	SELECT id, id_kendaraan, id_driver, created_at,created_by, updated_at, updated_by
	FROM driver_kendaraan
	WHERE deleted_at IS NULL
	ORDER BY id DESC`

	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var driver_kendaraan []model.DriverKendaraan
	for rows.Next() {
		var dk model.DriverKendaraan
		if err := rows.Scan(&dk.ID, &dk.ID_kendaraan, &dk.ID_driver, &dk.CreatedAt, &dk.CreatedBy, &dk.UpdatedAt, &dk.UpdatedBy); err != nil {
			return nil, err
		}
		driver_kendaraan = append(driver_kendaraan, dk)
	}
	return driver_kendaraan, nil

}

func (r *driver_kendaraanRepository) Update(id string, req model.UpdateDriverKendaraanRequest) (int64, error) {
	query := `UPDATE driver_kendaraan SET id_kendaraan=$1, id_driver=$2, updated_at=$3, updated_by=$4 WHERE id=$5 AND deleted_at IS NULL`
	result, err := database.DB.Exec(query, req.ID_kendaraan, req.ID_driver, time.Now(), req.UpdatedBy, id)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()

}

func (r *driver_kendaraanRepository) SoftDelete(id string, req model.DeleteDriverKendaraanRequest) (int64, error) {
	query := `UPDATE driver_kendaraan SET deleted_at=$1, deleted_by=$2 WHERE id=$3 AND deleted_at IS NULL`
	result, err := database.DB.Exec(query, time.Now(), req.DeletedBy, id)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}
