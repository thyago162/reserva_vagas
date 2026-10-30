package db

import (
	"reserva_vagas/internal/enderecos"
	"reserva_vagas/internal/escolaridades"
	"reserva_vagas/internal/escolas"
	"reserva_vagas/internal/turnos"
	vagasescolares "reserva_vagas/internal/vagas_escolares"

	"gorm.io/gorm"
)

var models = []any{
	&escolas.Escola{},
	&escolaridades.EtapaEnsino{},
	&escolaridades.Escolaridade{},
	&turnos.Turno{},
	&escolas.EscolaridadeEscola{},
	&enderecos.Endereco{},
	&vagasescolares.VagasEscolares{},
}

func RunMigrations(db *gorm.DB) error {
	return db.AutoMigrate(models...)
}
