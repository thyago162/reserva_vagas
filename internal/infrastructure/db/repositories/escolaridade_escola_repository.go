package repositories

import (
	"reserva_vagas/internal/escolas"

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
	tx := e.Db.Preload("escolaridade").Where("escola_id = ?", escolaID).Find(&escolaridadeEscolaLista)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return escolaridadeEscolaLista, nil
}

func (e *EscolaridadeEscolaRepositoryDb) Store(escolaridadeEscola *escolas.EscolaridadeEscola) error {
	tx := e.Db.Create(escolaridadeEscola)

	return tx.Error
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
