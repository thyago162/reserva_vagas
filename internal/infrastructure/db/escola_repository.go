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

func (e *EscolaRepositoryDb) GetById(id string) (*escolas.Escola, error) {
	var escola *escolas.Escola
	tx := e.Db.First(&escola, id)

	return escola, tx.Error
}

func (e *EscolaRepositoryDb) Update(escola *escolas.Escola) error {
	tx := e.Db.Save(escola)
	return tx.Error
}

func (e *EscolaRepositoryDb) Destroy(id string) error {
	tx := e.Db.Delete(&escolas.Escola{}, "id = ?", id)
	return tx.Error
}
