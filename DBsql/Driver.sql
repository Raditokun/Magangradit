CREATE TABLE driver (
    id SERIAL PRIMARY KEY,
    nama VARCHAR(100) NOT NULL,
    nip VARCHAR(50),
    foto TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR(100),
    updated_at TIMESTAMP,
    updated_by VARCHAR(100),
    deleted_at TIMESTAMP,
    deleted_by VARCHAR(100)
);