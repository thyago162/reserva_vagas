package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB() (*gorm.DB, error) {
	dsn := "reservas:reservas123.@tcp(127.0.0.1:3306)/reserva_vagas_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("falha ao conectar no banco: %w", err)
	}

	if err := RunMigrations(db); err != nil {
		return nil, fmt.Errorf("falha na migração: %w", err)
	}

	return db, nil
}
