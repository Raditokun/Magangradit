CREATE TABLE roles (
    id SERIAL PRIMARY KEY,
	role VARCHAR(20) NOT NULL CHECK (role IN ('admin', 'driver', 'user')),
    status VARCHAR(100) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR(100),
    updated_at TIMESTAMP,
	updated_by VARCHAR(100),
    deleted_at TIMESTAMP
);