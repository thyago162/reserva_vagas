package escolas

import (
	"reserva_vagas/internalerrors"

	"github.com/google/uuid"
)

type Escola struct {
	Id                 string               `gorm:"primaryKey;size:36"`
	Nome               string               `validate:"min=5,max=150,required"`
	Inep               string               `validate:"min=6,required"`
	Email              string               `validate:"email,required"`
	Sigla              string               `validate:"required,max=10"`
	EscolaridadeEscola []EscolaridadeEscola `gorm:"foreignKey:EscolaId"`
	EscolaTurno        []EscolaTurno        `gorm:"foreignKey:EscolaId"`
}

func NewEscola(nome string, inep string, email string, sigla string) (*Escola, error) {
	escola := &Escola{
		Id:    uuid.NewString(),
		Nome:  nome,
		Inep:  inep,
		Email: email,
		Sigla: sigla,
	}

	err := internalerrors.ValidateStruct(escola)
	if err == nil {
		return escola, nil
	}

	return nil, err
}
