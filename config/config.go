package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	var err error

	//initialize db

	db, err = InitializeSQLite()

	if err != nil {
		return fmt.Errorf("Error initializing db: %v", err)
	}

	return nil

}

func GetDB() *gorm.DB {
	return db
}

func GetLogger(prefix string) *Logger {
	logger = NewLogger(prefix)
	return logger
}
