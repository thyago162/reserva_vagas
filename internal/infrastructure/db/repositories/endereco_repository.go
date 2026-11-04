package repositories

import (
	"errors"
	"reserva_vagas/internal/enderecos"

	"gorm.io/gorm"
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

func (e *EnderecoRepositoryDb) StoreWithTx(tx any, endereco *enderecos.Endereco) error {
	gormTx, ok := tx.(*gorm.DB)
	if !ok {
		return errors.New("transação inválida ou tipo incorreto")
	}

	if err := gormTx.Create(endereco).Error; err != nil {
		return err
	}

	return nil
}

func (e *EnderecoRepositoryDb) Update(endereco *enderecos.Endereco) error {
	tx := e.Db.Save(endereco)
	return tx.Error
}

func (e *EnderecoRepositoryDb) Destroy(id string) error {
	tx := e.Db.Delete(&enderecos.Endereco{}, "id = ?", id)
	return tx.Error
}
