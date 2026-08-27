package seed

import (
	"log"
	"reserva_vagas/internal/turnos"

	"gorm.io/gorm"
)

func TurnoSeed(db *gorm.DB) {
	dados := []struct{ nome string }{
		{nome: "Integral"},
		{nome: "Manhã"},
		{nome: "Tarde"},
		{nome: "Noite"},
	}

	for _, d := range dados {
		var turnoExistente turnos.Turno

		err := db.Where("nome = ?", d.nome).First(&turnoExistente).Error
		if err != nil {
			novoTurno, err := turnos.NewTurno(d.nome)
			if err != nil {
				log.Printf("Erro ao validar seed de %s: %v", d.nome, err)
				continue
			}

			db.Create(novoTurno)
		}
	}
}
