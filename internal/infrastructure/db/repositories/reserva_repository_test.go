package repositories

import (
	"reserva_vagas/internal/importacao"
	"reserva_vagas/internal/reservas"
	"testing"

	"github.com/google/uuid"
)

func TestReservaRepository_CriarReserva(t *testing.T) {
	db := novoDBTeste(t)
	repo := NewReservaRepository(db)
	repoEscola := NewEscolaRepositoryDb(db)

	e1 := importacao.Escolaridade{Id: uuid.NewString(), Codigo: "FUND1"}
	turno := importacao.Turno{Id: uuid.NewString(), Codigo: "MANHA"}

	for _, v := range []any{&e1, &turno} {
		if err := db.Create(v).Error; err != nil {
			t.Fatalf("criar base: %v", err)
		}
	}

	escolaId := criarEscolaComVagas(t, db, "111111", e1.Id, turno.Id, 5)
	escola, err := repoEscola.Show(escolaId)
	if err != nil {
		t.Fatalf("buscar escola: %v", err)
	}

	if len(escola.EscolaridadeEscola) == 0 || len(escola.EscolaTurno) == 0 || len(escola.VagasEscolares) == 0 {
		t.Fatalf("escola sem associações carregadas: %+v", escola)
	}
	vaga := escola.VagasEscolares[0]

	candId := criarCandidato(t, db, "John F. Rambo", "52998224725", "23/09/2016",
		"Guile Nkifeout", "11144477735", "11/02/1984", "guile@email.com", "11999211010")

	reserva, err := reservas.NewReserva(escola.Id,
		e1.Id,
		turno.Id,
		vaga.Id,
		candId,
	)

	if err != nil {
		t.Fatalf("NewReserva: %v", err)
	}

	if err := repo.Store(reserva); err != nil {
		t.Fatalf("Store: %v", err)
	}

	// verifica persistência
	salva, err := repo.Show(reserva.Id)
	if err != nil {
		t.Fatalf("Show reserva: %v", err)
	}
	if salva.VagaId != vaga.Id {
		t.Errorf("VagaId = %s, quero %s", salva.VagaId, vaga.Id)
	}

	// verifica decremento da vaga
	depois, err := repoEscola.Show(escola.Id)
	if err != nil {
		t.Fatalf("Show escola: %v", err)
	}
	if got := depois.VagasEscolares[0].VagasTotais; got != 10 {
		t.Errorf("vagas restantes = %d, quero 9", got)
	}
}
