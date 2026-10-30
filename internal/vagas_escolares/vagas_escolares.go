package vagasescolares

import (
	"reserva_vagas/internalerrors"
	"time"

	"github.com/google/uuid"
)

type VagasEscolares struct {
	Id             string    `gorm:"primaryKey;size:36"`
	EscolaId       string    `gorm:"type:varchar(36);size:36;not null;index:idx_vagas_escolares_composicao,unique"`
	EscolaridadeId string    `gorm:"type:varchar(36);size:36;not null;index:idx_vagas_escolares_composicao,unique"`
	TurnoId        string    `gorm:"type:varchar(36);size:36;not null;index:idx_vagas_escolares_composicao,unique"`
	AnoLetivo      int32     `gorm:"type:int;not null;index:idx_vagas_escolares_composicao,unique"`
	Qtde           int32     `gorm:"type:int"`
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
