package repositories

import (
	"gorm.io/gorm"
	"reserva_vagas/internal/enderecos"
)

type EnderecoRepositoryDb struct {
	Db *gorm.DB
}

func NewEnderecoRepositorydb(db *gorm.DB) *EnderecoRepositoryDb {
	return &EnderecoRepositoryDb{Db: db}
}

func (e *EnderecoRepositoryDb) Store(endereco *enderecos.Endereco) error {
	tx := e.Db.Create(endereco)

	return tx.Error
}

func (e *EnderecoRepositoryDb) Update(endereco *enderecos.Endereco) error {
	tx := e.Db.Save(endereco)
	return tx.Error
}

func (e *EnderecoRepositoryDb) Destroy(id string) error {
	tx := e.Db.Delete(&enderecos.Endereco{}, "id = ?", id)
	return tx.Error
}