package seed

import (
	"log"
	"reserva_vagas/internal/escolaridades"

	"gorm.io/gorm"
)

func EtapaEnsinoSeed(db *gorm.DB) {
	dados := []struct {
		nome  string
		sigla string
	}{
		{nome: "Ensino Infantil", sigla: "EI"},
		{nome: "Ensino Fundamental", sigla: "EF"},
		{nome: "Ensino Médio", sigla: "EM"},
	}

	for _, d := range dados {
		var etapaExistente escolaridades.EtapaEnsino

		err := db.Where("sigla = ?", d.sigla).First(&etapaExistente).Error
		if err != nil {
			novaEtapa, err := escolaridades.NewEtapaEnsino(d.nome, d.sigla)
			if err != nil {
				log.Printf("Erro ao validar seed de %s: %v", d.nome, err)
				continue
			}

			db.Create(novaEtapa)
		}
	}
}
