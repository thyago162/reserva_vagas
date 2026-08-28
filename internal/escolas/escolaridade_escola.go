package escolas

import (
	"reserva_vagas/internal/escolaridades"
	"reserva_vagas/internalerrors"

	"github.com/google/uuid"
)

type EscolaridadeEscola struct {
	Id             string                      `gorm:"primaryKey"`
	EscolaId       string                      `gorm:"type:varchar(36);not null"`
	EscolaridadeId string                      `gorm:"type:varchar(36);not null"`
	Escola         *Escola                     `gorm:"foreignKey:EscolaId;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Escolaridade   *escolaridades.Escolaridade `gorm:"foreignKey:EscolaridadeId;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
}

func NewEscolaridadeEscola(escolaId string, escolaridadeId string) (*EscolaridadeEscola, error) {
	escolaridadeEscola := &EscolaridadeEscola{
		Id:             uuid.NewString(),
		EscolaId:       escolaId,
		EscolaridadeId: escolaridadeId,
	}

	err := internalerrors.ValidateStruct(escolaridadeEscola)
	if err != nil {
		return nil, err
	}

	return escolaridadeEscola, nil
}
