package config

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)


func SetupLogger() *os.File {
	
	if err := os.MkdirAll("logs", 0755); err != nil {
		log.Printf("Warning: Could not create logs directory: %v", err)
		return nil
	}

	
	logFileName := filepath.Join("logs", time.Now().Format("2006-01-02")+".log")
	logFile, err := os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Printf("Warning: Could not open log file: %v", err)
		return nil
	}

	
	multiWriter := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(multiWriter)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	return logFile
}
