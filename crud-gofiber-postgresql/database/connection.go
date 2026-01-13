package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {

	connStr := "host=localhost user=postgres password=040407 dbname=Silogis port=5432 sslmode=disable"

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Gagal connect DB:", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("Gagal ping DB:", err)
	}
	fmt.Println("Database SILOGIS Terkoneksi!")
}
