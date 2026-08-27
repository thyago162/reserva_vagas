package seed

import "gorm.io/gorm"

func Seed(db *gorm.DB) {

	EtapaEnsinoSeed(db)
	EscolaridadeSeed(db)
	TurnoSeed(db)
}
