package repository

import (
	"crud-app/app/model"
	"crud-app/database"
	"time"
)

//operation
//CRUD
type DriverRepository interface {
	GetAll() ([]model.Driver, error)
	Update(id string, req model.UpdateDriverRequest) (int64, error)
	SoftDelete(id string, req model.DeleteDriverRequest) (int64, error)
}

//method signature
type driverRepository struct{}

func NewDriverRepository() DriverRepository {
	return &driverRepository{}
}

//database schema
func (r *driverRepository) GetAll() ([]model.Driver, error) {
	query := `
		SELECT id, nama, nip, foto, created_at, created_by, updated_at, updated_by 
		FROM driver 
		WHERE deleted_at IS NULL
		ORDER BY id DESC
	`

	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var drivers []model.Driver
	for rows.Next() {
		var d model.Driver
		if err := rows.Scan(&d.ID, &d.Nama, &d.NIP, &d.Foto, &d.CreatedAt, &d.CreatedBy, &d.UpdatedAt, &d.UpdatedBy); err != nil {
			return nil, err
		}
		drivers = append(drivers, d)
	}

	return drivers, nil // ke model driver
}

func (r *driverRepository) Update(id string, req model.UpdateDriverRequest) (int64, error) {
	query := `UPDATE driver SET nama=$1, nip=$2, foto=$3, updated_at=$4, updated_by=$5 WHERE id=$6 AND deleted_at IS NULL`
	result, err := database.DB.Exec(query, req.Nama, req.NIP, req.Foto, time.Now(), req.UpdatedBy, id)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

func (r *driverRepository) SoftDelete(id string, req model.DeleteDriverRequest) (int64, error) {
	query := `UPDATE driver SET deleted_at=$1, deleted_by=$2 WHERE id=$3 AND deleted_at IS NULL`
	result, err := database.DB.Exec(query, time.Now(), req.DeletedBy, id)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}
