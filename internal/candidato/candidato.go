package candidato

import (
	"reserva_vagas/internalerrors"
	"time"

	"github.com/google/uuid"
)

type Candidato struct {
	Id                      string    `gorm:"primaryKey;size:36;column:id"`
	Nome                    string    `validate:"min=15,max=250,required" gorm:"not null"`
	Cpf                     string    `validate:"min=11,max=11,required" gorm:"not null"`
	DtNascimento            time.Time `validate:"required" gorm:"not null"`
	NomeResponsavel         string    `validate:"min=15,max=250,required" gorm:"not null"`
	CpfResponsavel          string    `validate:"min=11,max=11,required" gorm:"not null"`
	Telefone                string    `validate:"max=15"`
	DtNascimentoResponsavel time.Time `validate:"required" gorm:"not null"`
	EmailResponavel         string    `validate:"required" gorm:"not null"`
	CreatedAt               time.Time `gorm:"autoCreateTime"`
	UpdatedAt               time.Time `gorm:"autoCreateTime"`
}

type CreateCandidatoDTO struct {
	Nome                    string
	Cpf                     string
	DtNascimento            string
	NomeResponsavel         string
	CpfResponsavel          string
	Telefone                string
	DtNascimentoResponsavel string
	EmailResponavel         string
}

func NewCandidato(createCandidatoInput CreateCandidatoDTO) (*Candidato, error) {
	layoutBR := "02/01/2006"
	dtNascimento, err := time.Parse(layoutBR, createCandidatoInput.DtNascimento)
	if err != nil {
		return nil, err
	}

	dtNascimentoRes, err := time.Parse(layoutBR, createCandidatoInput.DtNascimentoResponsavel)
	if err != nil {
		return nil, err
	}

	candidato := &Candidato{
		Id:                      uuid.NewString(),
		Nome:                    createCandidatoInput.Nome,
		Cpf:                     createCandidatoInput.Cpf,
		DtNascimento:            dtNascimento,
		NomeResponsavel:         createCandidatoInput.NomeResponsavel,
		CpfResponsavel:          createCandidatoInput.CpfResponsavel,
		Telefone:                createCandidatoInput.Telefone,
		DtNascimentoResponsavel: dtNascimentoRes,
		EmailResponavel:         createCandidatoInput.EmailResponavel,
	}

	if err := internalerrors.ValidateStruct(candidato); err != nil {
		return nil, err
	}

	return candidato, nil
}
