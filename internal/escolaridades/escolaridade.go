package escolaridades

import (
	"fmt"
	"reserva_vagas/internalerrors"
	"time"

	"github.com/google/uuid"
)

type EtapaEnsino struct {
	Id        string    `gorm:"primaryKey;size:36;not null"`
	Nome      string    `validate:"required,max=50" gorm:"not null"`
	Sigla     string    `validate:"required,max=10" gorm:"uniqueIndex;type:varchar(10);not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type Escolaridade struct {
	Id            string       `gorm:"primaryKey;not null"`
	Nome          string       `validate:"required,max=50" gorm:"not null"`
	Codigo        string       `validate:"required" gorm:"not null"`
	CreatedAt     time.Time    `gorm:"autoCreateTime"`
	UpdatedAt     time.Time    `gorm:"autoUpdateTime"`
	EtapaEnsinoId string       `validate:"required" gorm:"type:varchar(36);not null"`
	EtapaEnsino   *EtapaEnsino `gorm:"foreignKey:EtapaEnsinoId;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
}

type RegraEscolaridade struct {
	Nome   string
	Idade  int
	Etapa  string
	Codigo string
}

var tabelaEscolaridades = []RegraEscolaridade{
	{Nome: "1 ano", Idade: 1, Etapa: "EI", Codigo: "EI01"},
	{Nome: "2 anos", Idade: 2, Etapa: "EI", Codigo: "EI02"},
	{Nome: "3 anos", Idade: 3, Etapa: "EI", Codigo: "EI03"},
	{Nome: "4 anos", Idade: 4, Etapa: "EI", Codigo: "EI04"},
	{Nome: "5 anos", Idade: 5, Etapa: "EI", Codigo: "EI05"},

	{Nome: "1º ano", Idade: 6, Etapa: "EF", Codigo: "EF01"},
	{Nome: "2º ano", Idade: 7, Etapa: "EF", Codigo: "EF02"},
	{Nome: "3º ano", Idade: 8, Etapa: "EF", Codigo: "EF03"},
	{Nome: "4º ano", Idade: 9, Etapa: "EF", Codigo: "EF04"},
	{Nome: "5º ano", Idade: 10, Etapa: "EF", Codigo: "EF05"},
	{Nome: "6º ano", Idade: 11, Etapa: "EF", Codigo: "EF06"},
	{Nome: "7º ano", Idade: 12, Etapa: "EF", Codigo: "EF07"},
	{Nome: "8º ano", Idade: 13, Etapa: "EF", Codigo: "EF08"},
	{Nome: "9º ano", Idade: 14, Etapa: "EF", Codigo: "EF019"},

	{Nome: "1º ano EM", Idade: 15, Etapa: "EM", Codigo: "EM01"},
	{Nome: "2º ano EM", Idade: 16, Etapa: "EM", Codigo: "EM02"},
	{Nome: "3º ano EM", Idade: 17, Etapa: "EM", Codigo: "EM03"},
}

func NewEtapaEnsino(nome, sigla string) (*EtapaEnsino, error) {
	etapaEnsino := &EtapaEnsino{
		Id:    uuid.New().String(),
		Nome:  nome,
		Sigla: sigla,
	}

	err := internalerrors.ValidateStruct(etapaEnsino)
	if err == nil {
		return etapaEnsino, nil
	}

	return nil, err
}

func NewEscolaridade(nome, etapaEnsinoId, codigo string) (*Escolaridade, error) {
	escolaridade := &Escolaridade{
		Id:            uuid.New().String(),
		Nome:          nome,
		EtapaEnsinoId: etapaEnsinoId,
		Codigo:        codigo,
	}

	err := internalerrors.ValidateStruct(escolaridade)
	if err == nil {
		return escolaridade, nil
	}

	return nil, err
}

func ObterEscolaridadePorIdade(idade int) (*RegraEscolaridade, error) {
	// Percorre o slice estático procurando a idade correspondente
	for _, regra := range tabelaEscolaridades {
		if regra.Idade == idade {
			return &regra, nil
		}
	}

	return nil, fmt.Errorf("nenhuma escolaridade padrão mapeada para a idade de %d anos", idade)
}
