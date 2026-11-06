package turnos

import (
	"reserva_vagas/internalerrors"

	"github.com/google/uuid"
)

type Turno struct {
	Id     string `gorm:"primaryKey;size:36;;not null"`
	Nome   string `validate:"required,min=5,max=10" gorm:"not null"`
	Codigo string `validate:"required" gorm:"not null"`
}

func NewTurno(nome, codigo string) (*Turno, error) {
	turno := &Turno{
		Id:     uuid.New().String(),
		Nome:   nome,
		Codigo: codigo,
	}

	err := internalerrors.ValidateStruct(turno)
	if err == nil {
		return turno, nil
	}

	return nil, err
}
