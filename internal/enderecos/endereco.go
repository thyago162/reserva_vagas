package enderecos

import (
	"reserva_vagas/internalerrors"
	"time"

	"github.com/google/uuid"
)

type Endereco struct {
	Id             string    `gorm:"primaryKey;size:36"`
	TipoLogradouro string    `validate:"required,max=50" gorm:"not null"`
	Logradouro     string    `validate:"required,max=150" gorm:"not null"`
	Numero         *string   `validate:"max=10" gorm:"default:'S/N'"`
	Complemento    string    `validate:"max=50"`
	Bairro         string    `validate:"required,max=50" gorm:"not null"`
	Cidade         string    `validate:"required,max=50" gorm:"not null"`
	Uf             string    `validate:"required,max=2" gorm:"not null"`
	Cep            string    `validate:"required,max=9" gorm:"not null"`
	Latitude       float64   `validate:"required" gorm:"not null"`
	Longitude      float64   `validate:"required" gorm:"not null"`
	EndereableID   string    `gorm:"index;not null"`
	EndereableType string    `gorm:"index;not null"`
	CreatedAt      time.Time `gorm:"autoCreateTime"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime"`
}

type CreateEnderecoInput struct {
	TipoLogradouro string
	Logradouro     string
	Numero         *string
	Complemento    string
	Bairro         string
	Cidade         string
	Uf             string
	Cep            string
	Latitude       float64
	Longitude      float64
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
