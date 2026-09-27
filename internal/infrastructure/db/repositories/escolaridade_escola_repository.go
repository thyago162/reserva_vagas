package repositories

import (
	"errors"
	"reserva_vagas/internal/escolas"
	"reserva_vagas/internalerrors"

	"gorm.io/gorm"
)

type EscolaridadeEscolaRepositoryDb struct {
	Db *gorm.DB
}

func NewEscolaridadeEscolaRepositoryDb(db *gorm.DB) *EscolaridadeEscolaRepositoryDb {
	return &EscolaridadeEscolaRepositoryDb{Db: db}
}

func (e *EscolaridadeEscolaRepositoryDb) Index(escolaID string) ([]*escolas.EscolaridadeEscola, error) {
	var escolaridadeEscolaLista []*escolas.EscolaridadeEscola
	tx := e.Db.Preload("Escolaridade").Where("escola_id = ?", escolaID).Find(&escolaridadeEscolaLista)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return escolaridadeEscolaLista, nil
}

func (e *EscolaridadeEscolaRepositoryDb) Store(ee []*escolas.EscolaridadeEscola) error {
	if len(ee) == 0 {
		return nil
	}

	if err := e.Db.Model(&escolas.EscolaridadeEscola{}).Create(ee).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return internalerrors.ErrRegistroExiste
		}
		return err
	}

	return nil
}

func (e *EscolaridadeEscolaRepositoryDb) Exists(escolaridadeID, escolaID string) (bool, error) {
	var id string

	tx := e.Db.Model(&escolas.EscolaridadeEscola{}).
		Select("id").
		Where("escola_id = ? AND escolaridade_id = ?", escolaID, escolaridadeID).
		Limit(1).
		Scan(&id)

	if tx.Error != nil {
		return false, tx.Error
	}

	return tx.RowsAffected > 0, nil
}

func (e *EscolaridadeEscolaRepositoryDb) Destroy(ids []string) error {
	if len(ids) == 0 {
		return errors.New("Não há nada para remover...")
	}

	tx := e.Db.Delete(&escolas.EscolaridadeEscola{}, "id IN ?", ids)
	if tx.RowsAffected == 0 {
		return errors.New("nenhum vínculo foi encontrado para os IDs informados")
	}

	return tx.Error
}
