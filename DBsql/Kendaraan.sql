CREATE TABLE kendaraan (
    id SERIAL PRIMARY KEY,
    id_jenis_kendaraan INT REFERENCES jenis_kendaraan(id), 
    nopol VARCHAR(50) UNIQUE NOT NULL,
    no_bpkb VARCHAR(50),
    no_mesin VARCHAR(50),
    no_rangka VARCHAR(50),
    nama_kendaraan VARCHAR(100),
    warna VARCHAR(50),
    kapasitas INT,
    deskripsi TEXT,
    status VARCHAR(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR(100),
    updated_at TIMESTAMP,
    updated_by VARCHAR(100),
    deleted_at TIMESTAMP,
    deleted_by VARCHAR(100)
);