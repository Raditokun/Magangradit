package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

func Connect(cfg DBConfig) error {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.Name,
		cfg.SSLMode,
	)

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("gagal connect DB: %w", err)
	}

	if err = DB.Ping(); err != nil {
		return fmt.Errorf("gagal ping DB: %w", err)
	}

	log.Println("Database SILOGIS Terkoneksi!")
	return nil
}

func Close() {
	if DB != nil {
		DB.Close()
	}
}
