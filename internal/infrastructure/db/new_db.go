package db

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB() (*gorm.DB, error) {

	dsn := os.Getenv("DATABASE")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("falha ao conectar no banco: %w", err)
	}

	if err := RunMigrations(db); err != nil {
		return nil, fmt.Errorf("falha na migração: %w", err)
	}

	return db, nil
}
