package main

import (
	"crud-app/config"
	"crud-app/database"
	"log"
	"fmt"
)

func main() {
	
	config.LoadEnv()

	
	logFile := config.SetupLogger()
	if logFile != nil {
		defer logFile.Close()
	}

	
	dbConfig := database.DBConfig{
		Host:     config.AppConfig.DBHost,
		Port:     config.AppConfig.DBPort,
		User:     config.AppConfig.DBUser,
		Password: config.AppConfig.DBPassword,
		Name:     config.AppConfig.DBName,
		SSLMode:  config.AppConfig.DBSSLMode,
	}

	if err := database.Connect(dbConfig); err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	
	app := config.NewApp()

	fmt.Printf("Server starting on port %s...", config.AppConfig.AppPort)
	log.Fatal(app.Listen(":" + config.AppConfig.AppPort))
}
