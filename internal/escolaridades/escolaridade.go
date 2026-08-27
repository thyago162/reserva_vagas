package escolaridades

import (
	"reserva_vagas/internalerrors"

	"github.com/google/uuid"
)

type EtapaEnsino struct {
	Id    string `gorm:"primaryKey"`
	Nome  string `validate:"min:3,required,max:15"`
	Sigla string `validate:"required,min:2,max:3" gorm:"uniqueIndex"`
}

type Escolaridade struct {
	Id            string      `gorm:"primaryKey"`
	Nome          string      `validate:"min:3,required,max:50"`
	EtapaEnsinoId string      `validate:"required" gorm:"type:varchar(36);not null"`
	EtapaEnsino   EtapaEnsino `gorm:"foreingKey:EtapaEnsinoId;constraint:CASCADE,OnDelete:RESTRICT;"`
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

func NewEscolaridade(nome, etapaEnsinoId string) (*Escolaridade, error) {
	escolaridade := &Escolaridade{
		Id:            uuid.New().String(),
		Nome:          nome,
		EtapaEnsinoId: etapaEnsinoId,
	}

	err := internalerrors.ValidateStruct(escolaridade)
	if err == nil {
		return escolaridade, nil
	}

	return nil, err
}
