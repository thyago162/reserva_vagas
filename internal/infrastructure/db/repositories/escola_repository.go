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
	tx := e.Db.
		Preload("Endereco").
		Preload("VagasEscolares").
		Find(&escolasLista)

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

func (e *EscolaRepositoryDb) GetByEscolaByEscolaridadeId(escolaridadeId string) ([]*escolas.Escola, error) {
	var escolasLista []*escolas.Escola

	tx := e.Db.
		Table("escolas").
		Select("escolas.id", "escolas.nome", "escolas.inep", "escolas.email", "escolas.sigla", "escolas.telefone").
		Joins("INNER JOIN escolaridade_escolas ON escolaridade_escolas.escola_id = escolas.id").
		Joins("INNER JOIN escolaridades ON escolaridade_escolas.escolaridade_id = escolaridades.id").
		Joins("INNER JOIN vagas_escolares ON vagas_escolares.escola_id = escolas.id AND vagas_escolares.escolaridade_id = escolaridades.id").
		Where("escolaridades.id = ? AND vagas_escolares.vagas_disponiveis > 0", escolaridadeId).
		Distinct("escolas.id").
		Preload("Endereco", func(db *gorm.DB) *gorm.DB {
			return db.Select(
				"id", "tipo_logradouro", "logradouro", "numero", "complemento",
				"bairro", "cidade", "uf", "cep", "endereable_id", "endereable_type",
				"latitude", "longitude",
			)
		}).
		Preload("VagasEscolares", func(db *gorm.DB) *gorm.DB {
			return db.
				Preload("Escolaridade").
				Preload("Turno").
				Where("escolaridade_id = ? AND vagas_disponiveis > 0", escolaridadeId).
				Select("id", "escola_id", "vagas_disponiveis", "escolaridade_id", "turno_id")
		}).
		Find(&escolasLista)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return escolasLista, nil
}
