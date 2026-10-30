package escolaridades

import (
	"reserva_vagas/internalerrors"
	"time"

	"github.com/google/uuid"
)

type EtapaEnsino struct {
	Id        string    `gorm:"primaryKey;size:36"`
	Nome      string    `validate:"required,max=50"`
	Sigla     string    `validate:"required,max=10" gorm:"uniqueIndex;type:varchar(10);not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type Escolaridade struct {
	Id            string       `gorm:"primaryKey"`
	Nome          string       `validate:"required,max=50"`
	Codigo        string       `validate:"required"`
	CreatedAt     time.Time    `gorm:"autoCreateTime"`
	UpdatedAt     time.Time    `gorm:"autoUpdateTime"`
	EtapaEnsinoId string       `validate:"required" gorm:"type:varchar(36);not null"`
	EtapaEnsino   *EtapaEnsino `gorm:"foreingKey:EtapaEnsinoId;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
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
