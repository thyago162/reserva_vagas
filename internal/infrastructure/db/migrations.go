package db

import (
	"reserva_vagas/internal/escolaridades"
	"reserva_vagas/internal/escolas"
	"reserva_vagas/internal/turnos"

	"gorm.io/gorm"
)

var models = []any{
	&escolas.Escola{},
	&escolaridades.EtapaEnsino{},
	&escolaridades.Escolaridade{},
	&turnos.Turno{},
}

func RunMigrations(db *gorm.DB) error {
	return db.AutoMigrate(models...)
}
