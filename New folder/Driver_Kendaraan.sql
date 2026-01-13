CREATE TABLE driver_kendaraan (
    id SERIAL PRIMARY KEY,
	id_kendaraan INT REFERENCES kendaraan(id)
	id_driver INT REFERENCES driver(id)
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR(100),
    updated_at TIMESTAMP,
    updated_by VARCHAR(100),
    deleted_at TIMESTAMP,
    deleted_by VARCHAR(100)
);