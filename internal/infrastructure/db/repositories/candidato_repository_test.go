package repositories

import (
	"reserva_vagas/internal/candidato"
	"testing"

	"gorm.io/gorm"
)

func criarCandidato(
	t *testing.T,
	db *gorm.DB,
	nomeCandidato,
	cpfCandidato,
	dtNascimentoCandidato,
	nomeResponsavel,
	cpfResponsavel,
	dtNascimentoResponsavel,
	email,
	telefone string,
) string {
	t.Helper()

	dto := candidato.CreateCandidatoDTO{
		Nome:                    nomeCandidato,
		Cpf:                     cpfCandidato,
		DtNascimento:            dtNascimentoCandidato,
		NomeResponsavel:         nomeResponsavel,
		CpfResponsavel:          cpfResponsavel,
		DtNascimentoResponsavel: dtNascimentoResponsavel,
		EmailResponavel:         email,
		Telefone:                telefone,
	}

	c, err := candidato.NewCandidato(dto)
	if err != nil {
		t.Fatalf("criar candidato: %v", err)
	}

	if err := db.Create(c).Error; err != nil {
		t.Fatalf("erro ao salvar: %v", err)
	}

	return c.Id
}
