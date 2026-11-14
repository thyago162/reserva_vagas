package repositories

import (
	"reserva_vagas/internal/vagasescolares"
	"reserva_vagas/internalerrors"

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

func (r *VagasEscolaresRepository) DebitarVaga(tx any, id string) error {
	gormTx, ok := tx.(*gorm.DB)

	if !ok {
		return internalerrors.ErrTransacaoInvalida
	}

	result := gormTx.Model(&vagasescolares.VagasEscolares{}).
		Where("id = ?", id).
		Where("vagas_disponiveis > 0").
		Updates(map[string]any{
			"vagas_disponiveis": gorm.Expr("vagas_disponiveis - ?", 1),
			"vagas_reservadas":  gorm.Expr("vagas_reservadas + ?", 1),
		})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return internalerrors.ErrVagasIndisponiveis
	}

	return nil
}
