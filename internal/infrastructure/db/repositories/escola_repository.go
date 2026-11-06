package repositories

import (
	"fmt"
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

func (e *EscolaRepositoryDb) GetByEscolaByCodigo(codigo string) ([]*escolas.Escola, error) {
	var escolasLista []*escolas.Escola
	fmt.Println(codigo)

	tx := e.Db.
		Distinct("escolas.id").
		// 1. Seleciona apenas o necessário da tabela principal
		Select("escolas.id", "escolas.nome", "escolas.inep", "escolas.email", "escolas.sigla", "escolas.telefone").

		// 2. CORREÇÃO: Faz o Join na query principal para FILTRAR as escolas corretas no banco
		Joins("JOIN escolaridade_escolas ON escolaridade_escolas.escola_id = escolas.id").
		Joins("JOIN escolaridades on escolaridade_escolas.escolaridade_id = escolaridade_id").
		Where("escolaridades.codigo = ?", codigo).

		// 3. Mantém o Preload para trazer os dados estruturados de Escolaridade (sem o Where aqui dentro)
		Preload("EscolaridadeEscola", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "escola_id", "escolaridade_id").
				Preload("Escolaridade", func(db *gorm.DB) *gorm.DB {
					return db.Select("id", "nome as nome_escolaridade")
				})
		}).
		// Seu mapeamento de endereço está perfeito (incluiu os campos polimórficos corretamente!)
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
		// Perfeito: incluiu o id e a FK escola_id
		Preload("VagasEscolares", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "escola_id", "vagas_disponiveis")
		}).
		// Perfeito: incluiu as chaves de ligação id, escola_id e turno_id
		Preload("EscolaTurno", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "escola_id", "turno_id").
				Preload("Turno", func(db *gorm.DB) *gorm.DB {
					return db.Select("id", "nome as nome_turno")
				})
		}).
		Find(&escolasLista)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return escolasLista, nil
}
