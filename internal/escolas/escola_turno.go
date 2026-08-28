package escolas

import (
	"reserva_vagas/internal/turnos"
	"reserva_vagas/internalerrors"

	"github.com/google/uuid"
)

type EscolaTurno struct {
	Id       string        `gorm:"primaryKey"`
	EscolaId string        `gorm:"type:varchar(36);not null"`
	TurnoId  string        `gorm:"type:varchar(36);not null"`
	Escola   *Escola       `gorm:"foreignKey:EscolaId;constraint:onUpdate:CASCADE,onDelete:RESTRICT"`
	Turno    *turnos.Turno `gorm:"foreignKey:TurnoId;constraint:onUpdate:CASCADE,onDelete:RESTRICT"`
}

func NewEscolaTurno(escolaId string, turnoId string) (*EscolaTurno, error) {
	escolaTurno := &EscolaTurno{
		Id:       uuid.NewString(),
		EscolaId: escolaId,
		TurnoId:  turnoId,
	}

	err := internalerrors.ValidateStruct(escolaTurno)
	if err != nil {
		return nil, err
	}

	return escolaTurno, nil
}
