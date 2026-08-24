package db

import (
	"fmt"
	"reserva_vagas/internal/escolas"

	"github.com/libtnb/sqlite"
	"gorm.io/gorm"
)

func NewDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("falha ao conectar no banco: %w", err)
	}

	if err := db.AutoMigrate(&escolas.Escola{}); err != nil {
		return nil, fmt.Errorf("falha na migração: %w", err)
	}

	return db, nil
}
