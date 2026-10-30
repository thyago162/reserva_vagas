package escolas

import (
	"reserva_vagas/internal/turnos"
	"reserva_vagas/internalerrors"
	"time"

	"github.com/google/uuid"
)

type EscolaTurno struct {
	Id        string        `gorm:"primaryKey;size:36"`
	EscolaId  string        `gorm:"type:varchar(36);size:36;not null;index:idx_escola_turno,unique"`
	TurnoId   string        `gorm:"type:varchar(36);size:36;not null;index:idx_escola_turno,unique"`
	CreatedAt time.Time     `gorm:"autoCreateTime"`
	UpdatedAt time.Time     `gorm:"autoUpdateTime"`
	Escola    *Escola       `gorm:"foreignKey:EscolaId;constraint:onUpdate:CASCADE,onDelete:RESTRICT"`
	Turno     *turnos.Turno `gorm:"foreignKey:TurnoId;constraint:onUpdate:CASCADE,onDelete:RESTRICT"`
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
