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
		Distinct("escolas.id").
		// 1. Seleciona apenas os campos necessários da tabela principal
		Select("escolas.id", "escolas.nome", "escolas.inep", "escolas.email", "escolas.sigla", "escolas.telefone").

		// 2. Faz o Join com a tabela pivot e com escolaridades para filtrar apenas as escolas que possuem essa escolaridade
		Joins("JOIN escolaridade_escolas ON escolaridade_escolas.escola_id = escolas.id").
		Joins("JOIN escolaridades ON escolaridade_escolas.escolaridade_id = escolaridades.id").
		Where("escolaridades.id = ?", escolaridadeId).

		// 3. Preload do Endereço (trazendo apenas os campos especificados)
		Preload("Endereco", func(db *gorm.DB) *gorm.DB {
			return db.Select(
				"id",
				"tipo_logradouro",
				"logradouro",
				"numero",
				"complemento",
				"bairro",
				"cidade",
				"uf",
				"cep",
				"endereable_id",
				"endereable_type",
				"latitude",
				"longitude",
			)
		}).

		// 4. Preload das Vagas Escolares (filtrando e trazendo os respectivos Turnos e Escolaridades)
		Preload("VagasEscolares", func(db *gorm.DB) *gorm.DB {
			return db.
				Preload("Escolaridade").
				Preload("Turno").
				Where("escolaridade_id = ?", escolaridadeId).
				Select("id", "escola_id", "vagas_disponiveis", "escolaridade_id", "turno_id")
		}).
		Find(&escolasLista)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return escolasLista, nil
}
