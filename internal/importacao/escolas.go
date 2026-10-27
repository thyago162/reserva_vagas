package importacao

import (
	"errors"
	"fmt"
	"mime/multipart"
	"reserva_vagas/internalerrors"

	"github.com/google/uuid"

	"github.com/gocarina/gocsv"
)

type Escola struct {
	Id       string `gorm:"primaryKey;size:36"`
	Nome     string
	Inep     string
	Email    string
	Sigla    string
	Telefone string
	Endereco Endereco `gorm:"polymorphic:Endereable;polymorphicValue:escola"`
}

type Endereco struct {
	Id             string `gorm:"primaryKey;size:36"`
	TipoLogradouro string
	Logradouro     string
	Numero         string
	Complemento    string
	Bairro         string
	Cidade         string
	Uf             string
	Cep            string
	Latitude       string
	Longitude      string

	EndereableID   string `gorm:"size:36"`
	EndereableType string
}

type ImportacaoEscolaDTO struct {
	Nome           string `csv:"nome" validate:"required,max=150"`
	Inep           string `csv:"inep" validate:"min=6,required"`
	Email          string `csv:"email" validate:"email,required"`
	Sigla          string `csv:"sigla" validate:"required,max=10"`
	Telefone       string `csv:"telefone" validate:"required,max=15"`
	TipoLogradouro string `csv:"tipo_logradouro" validate:"required,max=50"`
	Logradouro     string `csv:"logradouro" validate:"required,max=150"`
	Numero         string `csv:"numero" validate:"required,max=10"`
	Complemento    string `csv:"complemento" validate:"max=50"`
	Bairro         string `csv:"bairro" validate:"required,max=50"`
	Cidade         string `csv:"cidade" validate:"required,max=50"`
	Uf             string `csv:"uf" validate:"required,len=2"`
	Cep            string `csv:"cep" validate:"required,max=9"`
	Latitude       string `csv:"latitude" validate:"max=20"`
	Longitude      string `csv:"longitude" validate:"max=20"`
}

func NewImportacaoEscola(input ImportacaoEscolaDTO) (*Escola, error) {
	escolaID := uuid.New().String()
	enderecoID := uuid.New().String()
	if escolaID == "" {
		return nil, errors.New("falha ao gerar ID da escola")
	}

	importacaoEscola := &Escola{
		Id:       escolaID,
		Nome:     input.Nome,
		Inep:     input.Inep,
		Email:    input.Email,
		Sigla:    input.Sigla,
		Telefone: input.Telefone,
		Endereco: Endereco{
			Id:             enderecoID,
			EndereableID:   escolaID,
			EndereableType: "escola",
			TipoLogradouro: input.TipoLogradouro,
			Logradouro:     input.Logradouro,
			Numero:         input.Numero,
			Complemento:    input.Complemento,
			Bairro:         input.Bairro,
			Cidade:         input.Cidade,
			Uf:             input.Uf,
			Cep:            input.Cep,
			Latitude:       input.Latitude,
			Longitude:      input.Longitude,
		},
	}

	err := internalerrors.ValidateStruct(importacaoEscola)
	if err == nil {
		return importacaoEscola, nil
	}

	return nil, err
}

func (i *Escola) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"nome":            i.Nome,
		"inep":            i.Inep,
		"email":           i.Email,
		"sigla":           i.Sigla,
		"telefone":        i.Telefone,
		"tipo_logradouro": i.Endereco.TipoLogradouro,
		"logradouro":      i.Endereco.Logradouro,
		"numero":          i.Endereco.Numero,
		"complemento":     i.Endereco.Complemento,
		"bairro":          i.Endereco.Bairro,
		"cidade":          i.Endereco.Cidade,
		"uf":              i.Endereco.Uf,
		"cep":             i.Endereco.Cep,
		"latitude":        i.Endereco.Latitude,
		"longitude":       i.Endereco.Longitude,
	}
}

func LerCsvEscolas(file multipart.File) ([]Escola, error) {
	// 1. Lê o CSV usando a struct plana de DTO
	var linhasCsv []ImportacaoEscolaDTO
	if err := gocsv.Unmarshal(file, &linhasCsv); err != nil {
		return nil, err
	}

	var escolasParaSalvar []Escola

	// 2. Transforma cada linha do DTO na struct de domínio real, gerando os IDs
	for _, item := range linhasCsv {
		// Valida cada DTO individualmente se necessário
		if err := internalerrors.ValidateStruct(item); err != nil {
			return nil, err
		}

		escolaID := uuid.New().String()
		enderecoID := uuid.New().String()
		fmt.Println(item)

		escola := Escola{
			Id:       escolaID,
			Nome:     item.Nome,
			Inep:     item.Inep,
			Email:    item.Email,
			Sigla:    item.Sigla,
			Telefone: item.Telefone,
			Endereco: Endereco{
				Id:             enderecoID,
				TipoLogradouro: item.TipoLogradouro,
				Logradouro:     item.Logradouro,
				Numero:         item.Numero,
				Complemento:    item.Complemento,
				Bairro:         item.Bairro,
				Cidade:         item.Cidade,
				Uf:             item.Uf,
				Cep:            item.Cep,
				Latitude:       item.Latitude,
				Longitude:      item.Longitude,
				EndereableID:   escolaID,
				EndereableType: "escola",
			},
		}

		escolasParaSalvar = append(escolasParaSalvar, escola)
	}

	return escolasParaSalvar, nil
}
