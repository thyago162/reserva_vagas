package vagasescolares

import (
	"reserva_vagas/internalerrors"
	"time"

	"github.com/google/uuid"
)

type VagasEscolares struct {
	Id             string    `gorm:"primaryKey;size:36"`
	EscolaId       string    `gorm:"type:varchar(36);not null;uniqueIndex:idx_escola_id"`
	EscolaridadeId string    `gorm:"type:varchar(36);not null;uniqueIndex:idx_escolaridade_id"`
	TurnoId        string    `gorm:"type:varchar(36);not null;uniqueIndex:idx_turno_id"`
	Qtde           int32     `gorm:"type:int"`
	AnoLetivo      int32     `gorm:"type:int"`
	CreatedAt      time.Time `gorm:"autoCreateTime"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime"`
}

func NewVagasEscolares(escolaId, escolaridadeId, turnoId string, qtde, anoLetivo int32) (*VagasEscolares, error) {
	vagasEscolares := &VagasEscolares{
		Id:             uuid.NewString(),
		EscolaId:       escolaId,
		EscolaridadeId: escolaridadeId,
		TurnoId:        turnoId,
		Qtde:           qtde,
		AnoLetivo:      anoLetivo,
	}

	err := internalerrors.ValidateStruct(vagasEscolares)
	if err == nil {
		return vagasEscolares, nil
	}

	return nil, err
}
