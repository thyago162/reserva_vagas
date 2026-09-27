package repositories

import (
	"reserva_vagas/internal/escolaridades"

	"gorm.io/gorm"
)

type EscolaridadeRepositoryDb struct {
	Db *gorm.DB
}

func NewEscolaridadeRepositoryDb(db *gorm.DB) *EscolaridadeRepositoryDb {
	return &EscolaridadeRepositoryDb{Db: db}
}

func (e *EscolaridadeRepositoryDb) Show(id string) (*escolaridades.Escolaridade, error) {
	var escolaridade escolaridades.Escolaridade
	err := e.Db.Where("id = ?", id).First(&escolaridade).Error

	return &escolaridade, err
}
