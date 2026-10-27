package repositories

import (
	"reserva_vagas/internal/escolas"
	"reserva_vagas/internal/importacao"

	"gorm.io/gorm"
)

type EscolaRepositoryDb struct {
	Db *gorm.DB
}

func NewEscolaRepositoryDb(db *gorm.DB) *EscolaRepositoryDb {
	return &EscolaRepositoryDb{Db: db}
}

func (e *EscolaRepositoryDb) Store(escola *escolas.Escola) error {
	tx := e.Db.Create(escola)

	return tx.Error
}

func (e *EscolaRepositoryDb) Index() ([]*escolas.Escola, error) {
	var escolasLista []*escolas.Escola
	tx := e.Db.Preload("Endereco").Find(&escolasLista)

	return escolasLista, tx.Error
}

func (e *EscolaRepositoryDb) Show(id string) (*escolas.Escola, error) {
	var escola escolas.Escola
	err := e.Db.Preload("Endereco").Where("id = ?", id).First(&escola).Error

	return &escola, err
}

func (e *EscolaRepositoryDb) Update(escola *escolas.Escola) error {
	tx := e.Db.Save(escola)
	return tx.Error
}

func (e *EscolaRepositoryDb) Destroy(id string) error {
	tx := e.Db.Delete(&escolas.Escola{}, "id = ?", id)
	return tx.Error
}

func (e *EscolaRepositoryDb) ImportarEmLotes(escolas []importacao.Escola) error {
	if err := e.Db.CreateInBatches(&escolas, 100).Error; err != nil {
		return err
	}

	return nil
}
