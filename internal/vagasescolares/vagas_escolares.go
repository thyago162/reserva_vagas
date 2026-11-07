package vagasescolares

import (
	"reserva_vagas/internal/escolaridades"
	"reserva_vagas/internal/turnos"
	"reserva_vagas/internalerrors"
	"time"

	"github.com/google/uuid"
)

type VagasEscolares struct {
	Id               string                     `gorm:"primaryKey;size:36;not null"`
	EscolaId         string                     `gorm:"type:varchar(36);size:36;not null;index:idx_vagas_escolares_composicao,unique"`
	EscolaridadeId   string                     `gorm:"type:varchar(36);size:36;not null;index:idx_vagas_escolares_composicao,unique"`
	TurnoId          string                     `gorm:"type:varchar(36);size:36;not null;index:idx_vagas_escolares_composicao,unique"`
	AnoLetivo        int32                      `gorm:"type:int;not null;index:idx_vagas_escolares_composicao,unique"`
	VagasTotais      int32                      `gorm:"type:int;default:0"`
	VagasReservadas  int32                      `gorm:"type:int;default:0"`
	VagasDisponiveis int32                      `gorm:"type:int;default:0"`
	CreatedAt        time.Time                  `gorm:"autoCreateTime"`
	UpdatedAt        time.Time                  `gorm:"autoUpdateTime"`
	Turno            turnos.Turno               `gorm:"foreignKey:TurnoId"`
	Escolaridade     escolaridades.Escolaridade `gorm:"foreignKey:EscolaridadeId"`
}

func NewVagasEscolares(escolaId, escolaridadeId, turnoId string, vagasTotais, vagasReservadas, anoLetivo int32) (*VagasEscolares, error) {
	vagasEscolares := &VagasEscolares{
		Id:               uuid.NewString(),
		EscolaId:         escolaId,
		EscolaridadeId:   escolaridadeId,
		TurnoId:          turnoId,
		VagasTotais:      vagasTotais,
		VagasReservadas:  vagasReservadas,
		VagasDisponiveis: vagasTotais - vagasReservadas,
		AnoLetivo:        anoLetivo,
	}

	err := internalerrors.ValidateStruct(vagasEscolares)
	if err == nil {
		return vagasEscolares, nil
	}

	return nil, err
}
