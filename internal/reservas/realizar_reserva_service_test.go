package reservas_test

import (
	"reserva_vagas/internal/infrastructure/db/repositories"
	"reserva_vagas/internal/reservas"
	"reserva_vagas/internal/testutil"
	"sync"
	"testing"
)

type fakePublisher struct{}

func (f *fakePublisher) PublicarEmailReservaAsync(email string, nome string, senha string) {
}

const n = 5
const ttVagas = 4

func TestRealizarReserva_CriarVariasReservas(t *testing.T) {
	db := testutil.NovoDBTeste(t)
	reRepo := repositories.NewReservaRepository(db)
	caRepo := repositories.NewCandidatoRepository(db)
	enRepo := repositories.NewEnderecoRepositorydb(db)
	usRepo := repositories.NewUsuarioRepository(db)
	vaRepo := repositories.NewVagasEscolaresRepository(db)
	esRepo := repositories.NewEscolaRepositoryDb(db)

	s := reservas.NewRealizarReservaService(db, reRepo, caRepo, enRepo, usRepo, &fakePublisher{}, vaRepo)

	candidatos := []struct {
		Nome            string
		Cpf             string
		dtNascimento    string
		NomeReponsavel  string
		CpfResponsavel  string
		DtNascimentoRes string
		Email           string
		Telefone        string
	}{
		{"Candidato Um", "04212298819", "01/01/2025", "Responsavel Um", "02522343319", "12/12/2000", "res1@email.com", "1199991010"},
		{"Candidato Dois", "04212298818", "21/04/2025", "Responsavel Dois", "02522343318", "12/12/2002", "res2@email.com", "1199991011"},
		{"Candidato Tres", "04212298817", "12/02/2025", "Responsavel TrÊs", "02522343317", "12/12/1983", "res3@email.com", "1199991012"},
		{"Candidato Quatro", "04212298816", "03/03/2025", "Responsavel Quatro", "02522343316", "12/12/1995", "res4@email.com", "1199991013"},
		{"Candidato Cinco", "04212298815", "17/01/2025", "Responsavel Cinco", "02522343315", "12/12/1997", "res5@email.com", "1199991014"},
	}

	escola, err := testutil.CriarEscola(db, t, esRepo)
	if err != nil {
		t.Fatalf("criação escola: %v", err)
	}

	if len(escola.EscolaridadeEscola) == 0 || len(escola.EscolaTurno) == 0 || len(escola.VagasEscolares) == 0 {
		t.Fatalf("relacionamento incompleto: %v", err)
	}

	var wg sync.WaitGroup
	erros := make(chan error, n)
	sucess := make(chan string, n)
	numero := "100"

	for i := 0; i < n; i++ {
		wg.Add(1) //Isso aqui se não colocar um limite, não seria um problema?

		go func() {
			defer wg.Done() //Isso aqui é o que encerra a goroutine certo?
			candidato := reservas.CriarReservaDTO{
				NomeCandidato:           candidatos[i].Nome,
				CpfCandidato:            candidatos[i].Cpf,
				DtNascimentoCandidato:   candidatos[i].dtNascimento,
				NomeResponsavel:         candidatos[i].NomeReponsavel,
				DtNascimentoResponsavel: candidatos[i].DtNascimentoRes,
				CpfResponsavel:          candidatos[i].CpfResponsavel,
				EmailResponsavel:        candidatos[i].Email,
				Telefone:                candidatos[i].Telefone,
				EscolaId:                escola.Id,
				EscolaridadeId:          escola.EscolaridadeEscola[0].EscolaridadeId,
				TurnoId:                 escola.EscolaTurno[0].TurnoId,
				VagaId:                  escola.VagasEscolares[0].Id,
				TipoLogradouro:          "Rua",
				Logradouro:              "Rua do Teste",
				Numero:                  &numero,
				Complemento:             "Sem complemento",
				Bairro:                  "Bairro do teste",
				Cidade:                  "Cidade do teste",
				Uf:                      "XX",
				Cep:                     "58101100",
				Latitude:                -7.22222,
				Longitude:               -34.10000,
			}

			r, err := s.RealizarReserva(candidato)
			if err != nil {
				erros <- err
				return
			}

			if r.Situacao == "" {
				t.Errorf("Situação vazia para o candidato %s", candidato.NomeCandidato)
			}

			sucess <- r.CpfCandidato //A ideia é atribuir o cpf dos que tiveram sucesso
		}()
	}

	wg.Wait()
	close(erros)

	if len(sucess) != 4 || len(erros) != 1 {
		t.Errorf("Esperado %d com sucesso e %d com erro", ttVagas, 1)
	}

	vaga, err := vaRepo.Show(escola.VagasEscolares[0].Id)
	if err != nil {
		t.Errorf("erro: %v", err)
	}

	if vaga.VagasReservadas != 4 || vaga.VagasDisponiveis != 0 {
		t.Errorf("vagas reservadas = %d, esperado %d", vaga.VagasReservadas, ttVagas)
	}
}
