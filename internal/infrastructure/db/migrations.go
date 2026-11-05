package db

import (
	"reserva_vagas/internal/enderecos"
	"reserva_vagas/internal/escolaridades"
	"reserva_vagas/internal/escolas"
	"reserva_vagas/internal/reservas"
	"reserva_vagas/internal/turnos"
	"reserva_vagas/internal/usuario"
	"reserva_vagas/internal/vagasescolares"

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
	&escolas.EscolaTurno{},
	&reservas.Reserva{},
	&usuario.Usuario{},
}

func RunMigrations(db *gorm.DB) error {
	return db.AutoMigrate(models...)
}
