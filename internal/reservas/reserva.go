package reservas

import (
	"reserva_vagas/internal/candidato"
	"reserva_vagas/internal/escolaridades"
	"reserva_vagas/internal/escolas"
	"reserva_vagas/internal/turnos"
	"reserva_vagas/internal/vagasescolares"
	"reserva_vagas/internalerrors"
	"time"

	"github.com/google/uuid"
)

// 1. Defina um tipo customizado para a situação
type SituacaoReserva string

// 2. Crie as constantes válidas do Enum
const (
	SituacaoPendente   SituacaoReserva = "PENDENTE"
	SituacaoConfirmada SituacaoReserva = "CONFIRMADA"
	SituacaoCancelada  SituacaoReserva = "CANCELADA"
)

type Reserva struct {
	Id             string                         `gorm:"primaryKey;size:36;not null"`
	EscolaId       string                         `gorm:"size:36;not null"`
	EscolaridadeId string                         `gorm:"size:36;not null"`
	TurnoId        string                         `gorm:"size:36;not null"`
	VagaId         string                         `gorm:"size:36;not null"`
	CandidatoId    string                         `gorm:"size:36;not null"`
	Situacao       SituacaoReserva                `gorm:"type:enum('PENDENTE', 'CONFIRMADA', 'CANCELADA');default:'PENDENTE'"`
	CreatedAt      time.Time                      `gorm:"autoCreateTime"`
	UpdatedAt      time.Time                      `gorm:"autoCreateTime"`
	Escola         *escolas.Escola                `gorm:"foreignKey:EscolaId"`
	Escolaridade   *escolaridades.Escolaridade    `gorm:"foreignKey:EscolaridadeId"`
	Turno          *turnos.Turno                  `gorm:"foreignKey:TurnoId"`
	Vaga           *vagasescolares.VagasEscolares `gorm:"foreignKey:VagaId"`
	Candidato      *candidato.Candidato           `gorm:"foreignKey:CandidatoId"`
}

func NewReserva(escolaId, escolaridadeId, turnoId, vagaId, candidatoId string) (*Reserva, error) {
	reserva := &Reserva{
		Id:             uuid.NewString(),
		EscolaId:       escolaId,
		EscolaridadeId: escolaridadeId,
		TurnoId:        turnoId,
		VagaId:         vagaId,
		CandidatoId:    candidatoId,
		Situacao:       SituacaoPendente,
	}

	err := internalerrors.ValidateStruct(reserva)
	if err != nil {
		return nil, err
	}

	return reserva, nil
}
