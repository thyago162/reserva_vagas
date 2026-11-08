package seed

import (
	"log"
	"reserva_vagas/internal/turnos"

	"gorm.io/gorm"
)

func TurnoSeed(db *gorm.DB) {
	dados := []struct{ nome, codigo string }{
		{nome: "Integral", codigo: "IN"},
		{nome: "Manhã", codigo: "MN"},
		{nome: "Tarde", codigo: "TD"},
	}

	for _, d := range dados {
		var turnoExistente turnos.Turno

		err := db.Where("nome = ?", d.nome).First(&turnoExistente).Error
		if err != nil {
			novoTurno, err := turnos.NewTurno(d.nome, d.codigo)
			if err != nil {
				log.Printf("Erro ao validar seed de %s: %v", d.nome, err)
				continue
			}

			db.Create(novoTurno)
		}
	}
}
