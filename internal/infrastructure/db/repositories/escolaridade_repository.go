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

func (e *EscolaridadeRepositoryDb) FindByIds(ids []string) (int64, error) {
	var count int64
	err := e.Db.Model(&escolaridades.Escolaridade{}).Where("id in ?", ids).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}
