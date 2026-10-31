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

func (v *VagasEscolaresRepository) Index(escolaId string, anoLetivo int) ([]*vagasescolares.VagasEscolares, error) {
	var vagasEscolares []*vagasescolares.VagasEscolares

	tx := v.Db.Where("escola_id = ? AND ano_letivo = ", escolaId, anoLetivo).Find(&vagasEscolares)

	return vagasEscolares, tx.Error
}

func (v *VagasEscolaresRepository) Store(vagasEscolares []*vagasescolares.VagasEscolares) error {
	tx := v.Db.Create(&vagasEscolares)

	return tx.Error
}

func (v *VagasEscolaresRepository) Update(vagasescolares *vagasescolares.VagasEscolares) error {
	tx := v.Db.Save(&vagasescolares)

	return tx.Error
}

func (v *VagasEscolaresRepository) Show(id string) (*vagasescolares.VagasEscolares, error) {
	var vagaEscola vagasescolares.VagasEscolares

	err := v.Db.Where("id = ?", id).First(&vagaEscola).Error
	if err != nil {
		return nil, err
	}

	return &vagaEscola, err
}
