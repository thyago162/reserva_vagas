package candidato

import (
	"reserva_vagas/internalerrors"
	"time"
)

type Candidato struct {
	Id                      string    `gorm:"primaryKey;size:36"`
	Nome                    string    `validate:"min=15,max=250,required"`
	Cpf                     string    `validate:"min=11,max=11,required"`
	DtNascimento            string    `validate:"required"`
	NomeResponsavel         string    `validate:"min=15,max=250,required"`
	CpfResponsavel          string    `validate:"min=11,max=11,required"`
	Telefone                string    `validate:"required,max=15"`
	DtNascimentoResponsavel string    `validate:"required"`
	EmailResponavel         string    `validate:"required"`
	CreatedAt               time.Time `gorm:"autoCreateTime"`
	UpdatedAt               time.Time `gorm:"autoCreateTime"`
}

type CreateCandidatoInput struct {
	Nome                    string
	Cpf                     string
	DtNascimento            string
	NomeResponsavel         string
	CpfResponsavel          string
	Telefone                string
	DtNascimentoResponsavel string
	EmailResponavel         string
}

func NewCandidato(createCandidatoInput CreateCandidatoInput) (*Candidato, error) {
	candidato := &Candidato{
		Nome:                    createCandidatoInput.Nome,
		Cpf:                     createCandidatoInput.Cpf,
		DtNascimento:            createCandidatoInput.DtNascimento,
		NomeResponsavel:         createCandidatoInput.NomeResponsavel,
		CpfResponsavel:          createCandidatoInput.CpfResponsavel,
		Telefone:                createCandidatoInput.Telefone,
		DtNascimentoResponsavel: createCandidatoInput.DtNascimentoResponsavel,
		EmailResponavel:         createCandidatoInput.EmailResponavel,
	}

	if err := internalerrors.ValidateStruct(candidato); err != nil {
		return nil, err
	}

	return candidato, nil
}
