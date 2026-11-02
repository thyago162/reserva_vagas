package reservas

import (
	"reserva_vagas/internal/candidato"
	"reserva_vagas/internal/escolaridades"
	"reserva_vagas/internal/escolas"
	"reserva_vagas/internal/turnos"
	"reserva_vagas/internal/vagasescolares"
	"reserva_vagas/internalerrors"
	"time"
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
	Id             string                         `gorm:"primaryKey;size:36"`
	EscolaId       string                         `gorm:"size:36"`
	EscolaridadeId string                         `gorm:"size:36"`
	TurnoId        string                         `gorm:"size:36"`
	VagaId         string                         `gorm:"size:36"`
	CandidatoId    string                         `gorm:"size:36"`
	Situacao       SituacaoReserva                `gorm:"type:enum('PENDENTE', 'CONFIRMADA', 'CANCELADA');default:'PENDENTE'"`
	CreatedAt      time.Time                      `gorm:"autoCreateTime"`
	UpdatedAt      time.Time                      `gorm:"autoCreateTime"`
	Escola         *escolas.Escola                `gorm:"foreignKey:EscolaId"`
	Escolaridade   *escolaridades.Escolaridade    `gorm:"foreignKey:EscolaridadeId"`
	Turno          *turnos.Turno                  `gorm:"foreignKey:TurnoId"`
	Vaga           *vagasescolares.VagasEscolares `gorm:"foreignKey:VagaId"`
	Candidato      *candidato.Candidato           `gorm:"foreignKey:CandidatoId"`
}

func NewReserva(escolaId, escolaridadeId, turnoId, vagaId, candidatoId string) *Reserva {
	reserva := &Reserva{
		EscolaId:       escolaId,
		EscolaridadeId: escolaridadeId,
		TurnoId:        turnoId,
		VagaId:         vagaId,
		CandidatoId:    candidatoId,
		Situacao:       SituacaoPendente,
	}

	err := internalerrors.ValidateStruct(reserva)
	if err != nil {
		return nil
	}

	return reserva
}
