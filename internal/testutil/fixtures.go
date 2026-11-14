package testutil

import (
	"reserva_vagas/internal/candidato"
	"reserva_vagas/internal/escolas"
	"reserva_vagas/internal/importacao"
	"testing"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CriarEscola(db *gorm.DB, t *testing.T, repoEscola escolas.Repository) (*escolas.Escola, error) {
	e1 := importacao.Escolaridade{Id: uuid.NewString(), Codigo: "FUND1"}
	turno := importacao.Turno{Id: uuid.NewString(), Codigo: "MANHA"}

	for _, v := range []any{&e1, &turno} {
		if err := db.Create(v).Error; err != nil {
			t.Fatalf("criar base: %v", err)
		}
	}

	escolaId := CriarEscolaComVagas(t, db, "111111", e1.Id, turno.Id, 4)
	escola, err := repoEscola.Show(escolaId)
	if err != nil {
		t.Fatalf("criar escola: %v", err)
	}

	return escola, nil
}

// criarEscolaComVagas insere uma escola com vagas para uma escolaridade e devolve o id.
func CriarEscolaComVagas(t *testing.T, db *gorm.DB, inep, escolaridadeID, turnoID string, disponiveis int32) string {
	t.Helper()

	escolaID := uuid.NewString()
	escola := importacao.Escola{
		Id: escolaID, Nome: "Escola " + inep, Inep: inep,
		Email: inep + "@escola.test", Sigla: "E" + inep, Telefone: "83999990000",
		Endereco: importacao.Endereco{
			Id: uuid.NewString(), EndereableID: escolaID, EndereableType: "escola",
			Cidade: "João Pessoa", Uf: "PB", Latitude: -7.11, Longitude: -34.86, TipoLogradouro: "Rua", Logradouro: "Rua do Teste", Bairro: "Jardim Oc", Cep: "58102-121",
		},
	}
	if err := db.Omit("Endereco").Create(&escola).Error; err != nil {
		t.Fatalf("criar escola: %v", err)
	}

	pivoE := importacao.EscolaridadeEscola{Id: uuid.NewString(), EscolaId: escolaID, EscolaridadeId: escolaridadeID}
	if err := db.Create(&pivoE).Error; err != nil {
		t.Fatalf("criar pivô: %v", err)
	}

	pivoT := importacao.EscolaTurno{Id: uuid.NewString(), EscolaId: escolaID, TurnoId: turnoID}
	if err := db.Create(&pivoT).Error; err != nil {
		t.Fatalf("criar pivô: %v", err)
	}

	vaga := importacao.VagasEscolares{
		Id: uuid.NewString(), EscolaId: escolaID, EscolaridadeId: escolaridadeID, TurnoId: turnoID,
		VagasTotais: disponiveis, VagasDisponiveis: disponiveis, AnoLetivo: 2026,
	}
	if err := db.Create(&vaga).Error; err != nil {
		t.Fatalf("criar vaga: %v", err)
	}

	return escolaID
}

func CriarCandidato(
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
