package repositories

import (
	"errors"
	"testing"

	"reserva_vagas/internal/escolas"
	"reserva_vagas/internal/importacao"

	"github.com/glebarez/sqlite"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// ---------- Auxiliares ----------

// novoDBTeste cria um banco em memória limpo para cada teste.
func novoDBTeste(t *testing.T) *gorm.DB {
	t.Helper()

	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent), // use .Debug() no teste para ver o SQL
	})
	if err != nil {
		t.Fatalf("abrir banco: %v", err)
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(1) // em memória, cada conexão seria um banco diferente
	t.Cleanup(func() { sqlDB.Close() })

	err = db.AutoMigrate(
		&importacao.Escola{},
		&importacao.Endereco{},
		&importacao.Escolaridade{},
		&importacao.Turno{},
		&importacao.EscolaridadeEscola{},
		&importacao.EscolaTurno{},
		&importacao.VagasEscolares{},
	)
	if err != nil {
		t.Fatalf("migrar tabelas: %v", err)
	}

	return db
}

// criarEscolaComVagas insere uma escola com vagas para uma escolaridade e devolve o id.
func criarEscolaComVagas(t *testing.T, db *gorm.DB, inep, escolaridadeID, turnoID string, disponiveis int32) string {
	t.Helper()

	escolaID := uuid.NewString()
	escola := importacao.Escola{
		Id: escolaID, Nome: "Escola " + inep, Inep: inep,
		Email: inep + "@escola.test", Sigla: "E" + inep, Telefone: "83999990000",
		Endereco: importacao.Endereco{
			Id: uuid.NewString(), EndereableID: escolaID, EndereableType: "escola",
			Cidade: "João Pessoa", Uf: "PB", Latitude: -7.11, Longitude: -34.86,
		},
	}
	if err := db.Create(&escola).Error; err != nil {
		t.Fatalf("criar escola: %v", err)
	}

	pivo := importacao.EscolaridadeEscola{Id: uuid.NewString(), EscolaId: escolaID, EscolaridadeId: escolaridadeID}
	if err := db.Create(&pivo).Error; err != nil {
		t.Fatalf("criar pivô: %v", err)
	}

	vaga := importacao.VagasEscolares{
		Id: uuid.NewString(), EscolaId: escolaID, EscolaridadeId: escolaridadeID, TurnoId: turnoID,
		VagasTotais: 10, VagasDisponiveis: disponiveis, AnoLetivo: 2026,
	}
	if err := db.Create(&vaga).Error; err != nil {
		t.Fatalf("criar vaga: %v", err)
	}

	return escolaID
}

// ---------- Testes ----------

func TestEscolaRepository_StoreEShow(t *testing.T) {
	db := novoDBTeste(t)
	repo := NewEscolaRepositoryDb(db)

	escola := &escolas.Escola{
		Id: uuid.NewString(), Nome: "Escola Modelo", Inep: "123456",
		Email: "modelo@escola.test", Sigla: "EM", Telefone: "83999990000",
	}
	if err := repo.Store(escola); err != nil {
		t.Fatalf("Store: %v", err)
	}

	got, err := repo.Show(escola.Id)
	if err != nil {
		t.Fatalf("Show: %v", err)
	}
	if got.Nome != escola.Nome {
		t.Errorf("Nome = %q, esperado %q", got.Nome, escola.Nome)
	}
}

// Deve FALHAR hoje: Show devolve &escola (não nil) mesmo com erro.
func TestEscolaRepository_Show_NaoEncontrada(t *testing.T) {
	repo := NewEscolaRepositoryDb(novoDBTeste(t))

	got, err := repo.Show(uuid.NewString())

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		t.Errorf("erro = %v, esperado ErrRecordNotFound", err)
	}
	if got != nil {
		t.Errorf("esperado nil quando não encontra, veio %+v", got)
	}
}

// Deve FALHAR hoje: Save faz upsert e cria a escola que não existia.
func TestEscolaRepository_Update_NaoCriaEscolaInexistente(t *testing.T) {
	db := novoDBTeste(t)
	repo := NewEscolaRepositoryDb(db)

	err := repo.Update(&escolas.Escola{Id: uuid.NewString(), Nome: "Fantasma"})

	var total int64
	db.Model(&importacao.Escola{}).Count(&total)
	if total != 0 {
		t.Errorf("Update criou %d escola(s); não deveria criar nada", total)
	}
	if err == nil {
		t.Error("esperado erro ao atualizar escola inexistente")
	}
}

// Deve FALHAR hoje se o Distinct sobrescrever o Select (campos como Nome vêm vazios).
func TestEscolaRepository_GetByEscolaByEscolaridadeId(t *testing.T) {
	db := novoDBTeste(t)
	repo := NewEscolaRepositoryDb(db)

	e1 := importacao.Escolaridade{Id: uuid.NewString(), Codigo: "FUND1"}
	e2 := importacao.Escolaridade{Id: uuid.NewString(), Codigo: "FUND2"}
	turno := importacao.Turno{Id: uuid.NewString(), Codigo: "MANHA"}
	for _, v := range []any{&e1, &e2, &turno} {
		if err := db.Create(v).Error; err != nil {
			t.Fatalf("criar base: %v", err)
		}
	}

	comVaga := criarEscolaComVagas(t, db, "111111", e1.Id, turno.Id, 5)
	criarEscolaComVagas(t, db, "222222", e1.Id, turno.Id, 0) // sem vagas disponíveis
	criarEscolaComVagas(t, db, "333333", e2.Id, turno.Id, 5) // outra escolaridade

	got, err := repo.GetByEscolaByEscolaridadeId(e1.Id)
	if err != nil {
		t.Fatalf("consulta: %v", err)
	}

	if len(got) != 1 {
		t.Fatalf("escolas = %d, esperado 1 (só a com vagas na escolaridade)", len(got))
	}
	if got[0].Id != comVaga {
		t.Errorf("escola errada: %s", got[0].Id)
	}
	if got[0].Nome == "" {
		t.Error("Nome vazio: o Select foi sobrescrito (veja o Distinct)")
	}
	if len(got[0].VagasEscolares) != 1 {
		t.Errorf("vagas carregadas = %d, esperado 1", len(got[0].VagasEscolares))
	}
}
