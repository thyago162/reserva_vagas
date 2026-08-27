package turnos

import (
	"reserva_vagas/internalerrors"

	"github.com/google/uuid"
)

type Turno struct {
	Id   string `gorm:"primaryKey"`
	Nome string `validate:"required,min=5,max=10"`
}

func NewTurno(nome string) (*Turno, error) {
	turno := &Turno{Id: uuid.New().String(), Nome: nome}

	err := internalerrors.ValidateStruct(turno)
	if err == nil {
		return turno, nil
	}

	return nil, err
}
