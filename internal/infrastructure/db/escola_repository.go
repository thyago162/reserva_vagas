package db

import (
	"reserva_vagas/internal/escolas"

	"gorm.io/gorm"
)

type EscolaRepositoryDb struct {
	Db *gorm.DB
}

func NewEscolaRepositoryDb(db *gorm.DB) *EscolaRepositoryDb {
	return &EscolaRepositoryDb{Db: db}
}

func (e *EscolaRepositoryDb) Salvar(escola *escolas.Escola) error {
	tx := e.Db.Create(escola)

	return tx.Error
}

func (e *EscolaRepositoryDb) GetAll() ([]*escolas.Escola, error) {
	var escolasLista []*escolas.Escola
	tx := e.Db.Find(&escolasLista)

	return escolasLista, tx.Error
}
