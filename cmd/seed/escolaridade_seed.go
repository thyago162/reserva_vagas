package seed

import (
	"fmt"
	"log"
	"reserva_vagas/internal/escolaridades"

	"gorm.io/gorm"
)

func EscolaridadeSeed(db *gorm.DB) {
	dados := []struct {
		nome       string
		siglaEtapa string
	}{
		{nome: "1 ano", siglaEtapa: "EI"},
		{nome: "2 anos", siglaEtapa: "EI"},
		{nome: "3 anos", siglaEtapa: "EI"},
		{nome: "4 anos", siglaEtapa: "EI"},
		{nome: "5 anos", siglaEtapa: "EI"},
		{nome: "1º ano", siglaEtapa: "EF"},
		{nome: "2º ano", siglaEtapa: "EF"},
		{nome: "3º ano", siglaEtapa: "Ef"},
		{nome: "4º ano", siglaEtapa: "EF"},
		{nome: "5º ano", siglaEtapa: "EF"},
		{nome: "6º ano", siglaEtapa: "EF"},
		{nome: "7º ano", siglaEtapa: "EF"},
		{nome: "8º ano", siglaEtapa: "EF"},
		{nome: "9º ano", siglaEtapa: "EF"},
		{nome: "1º ano", siglaEtapa: "EM"},
		{nome: "2º ano", siglaEtapa: "EM"},
		{nome: "3º ano", siglaEtapa: "EM"},
	}

	for _, d := range dados {
		var etapaEnsino escolaridades.EtapaEnsino
		err := db.Where("sigla = ?", d.siglaEtapa).First(&etapaEnsino)
		if err != nil {
			fmt.Printf("Registro não encontrado")
			continue
		}

		var escolaridadeExistente escolaridades.Escolaridade
		err := db.Where("nome = ? and etapa_ensino_id = ?", d.nome, etapaEnsino.Id).First(&escolaridadeExistente).Error

		if err != nil {
			novaEscolaridade, err := escolaridades.NewEscolaridade(d.nome, etapaEnsino.Id)
			if err != nil {
				log.Printf("Erro ao validar seed de %s (%s): %v", d.nome, d.siglaEtapa, err)
				continue
			}

			db.Create(novaEscolaridade)
		}

	}
}
