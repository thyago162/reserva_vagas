package seed

import (
	"log"
	"reserva_vagas/internal/escolaridades"

	"gorm.io/gorm"
)

func EscolaridadeSeed(db *gorm.DB) {
	dados := []struct {
		nome       string
		siglaEtapa string
		codigo     string
	}{
		// Educação Infantil
		{nome: "1 ano", siglaEtapa: "EI", codigo: "EI01"},
		{nome: "2 anos", siglaEtapa: "EI", codigo: "EI02"},
		{nome: "3 anos", siglaEtapa: "EI", codigo: "EI03"},
		{nome: "4 anos", siglaEtapa: "EI", codigo: "EI04"},
		{nome: "5 anos", siglaEtapa: "EI", codigo: "EI05"},

		// Ensino Fundamental (Siglas corrigidas para "EF")
		{nome: "1º ano", siglaEtapa: "EF", codigo: "EF01"},
		{nome: "2º ano", siglaEtapa: "EF", codigo: "EF02"},
		{nome: "3º ano", siglaEtapa: "EF", codigo: "EF03"},
		{nome: "4º ano", siglaEtapa: "EF", codigo: "EF04"},
		{nome: "5º ano", siglaEtapa: "EF", codigo: "EF05"},
		{nome: "6º ano", siglaEtapa: "EF", codigo: "EF06"},
		{nome: "7º ano", siglaEtapa: "EF", codigo: "EF07"},
		{nome: "8º ano", siglaEtapa: "EF", codigo: "EF08"},
		{nome: "9º ano", siglaEtapa: "EF", codigo: "EF09"},
	}

	for _, d := range dados {
		// 1. Busca a Etapa de Ensino correspondente
		var etapaEnsino escolaridades.EtapaEnsino
		err := db.Where("sigla = ?", d.siglaEtapa).First(&etapaEnsino).Error
		if err != nil {
			log.Printf("Etapa de ensino com sigla '%s' não encontrada para a escolaridade '%s'", d.siglaEtapa, d.nome)
			continue
		}

		// 2. Checa se essa escolaridade JÁ EXISTE para ESTA ETAPA ESPECÍFICA
		var escolaridadeExistente escolaridades.Escolaridade
		err = db.Where("nome = ? AND etapa_ensino_id = ? AND codigo = ?", d.nome, etapaEnsino.Id, d.codigo).First(&escolaridadeExistente).Error

		// Se NÃO encontrou (err != nil), podemos criar
		if err != nil {
			novaEscolaridade, err := escolaridades.NewEscolaridade(d.nome, etapaEnsino.Id, d.codigo)
			if err != nil {
				log.Printf("Erro ao validar seed de %s (%s): %v", d.nome, d.siglaEtapa, err)
				continue
			}

			db.Create(novaEscolaridade)
		}
	}
}
