package repositories

import (
	vagasescolares "reserva_vagas/internal/vagas_escolares"

	"gorm.io/gorm"
)

type VagasEscolaresRepository struct {
	Db *gorm.DB
}

func NewVagasEscolaresRepository(db *gorm.DB) *VagasEscolaresRepository {
	return &VagasEscolaresRepository{Db: db}
}

func (v *VagasEscolaresRepository) Index(escolaId string) ([]*vagasescolares.VagasEscolares, error) {
	var vagasEscolares []*vagasescolares.VagasEscolares

	tx := v.Db.Where("escola_id = ?", escolaId).Find(&vagasEscolares)

	return vagasEscolares, tx.Error
}

func (v *VagasEscolaresRepository) Store(vagasEscolares []*vagasescolares.VagasEscolares) error {
	tx := v.Db.CreateInBatches(&vagasEscolares, 10)

	return tx.Error
}
