package enderecos

import (
	"reserva_vagas/internalerrors"
	"time"

	"github.com/google/uuid"
)

type Endereco struct {
	Id             string    `gorm:"primaryKey;size:36"`
	TipoLogradouro string    `validate:"required,max=50"`
	Logradouro     string    `validate:"required,max=150"`
	Numero         string    `validate:"required,max=10"`
	Complemento    string    `validate:"max=50"`
	Bairro         string    `validate:"required,max=50"`
	Cidade         string    `validate:"required,max=50"`
	Uf             string    `validate:"required,max=2"`
	Cep            string    `validate:"required,max=9"`
	Latitude       string    `validate:"max=20"`
	Longitude      string    `validate:"max=20"`
	EndereableID   string 	 `gorm:"index"`
    EndereableType string 	 `gorm:"index"`
	CreatedAt      time.Time `gorm:"autoCreateTime"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime"`
}

type CreateEnderecoInput struct {
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
	EndereableID   string
	EndereableType string
}

func NewEndereco(input CreateEnderecoInput) (*Endereco, error) {
	endereco := &Endereco{
		Id:             uuid.New().String(),
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
		EndereableID:   input.EndereableID,
		EndereableType: input.EndereableType,
	}

	err := internalerrors.ValidateStruct(endereco)
	if err == nil {
		return endereco, nil
	}

	return nil, err
}
