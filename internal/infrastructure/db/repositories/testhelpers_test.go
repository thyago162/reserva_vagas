package repositories

import (
	"os"
	"reserva_vagas/internal/importacao"
	"testing"

	"github.com/glebarez/sqlite"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func novoDBTeste(t *testing.T) *gorm.DB {
	t.Helper()
	_ = godotenv.Load("../../../../.env.test")
	dsn := os.Getenv("TEST_MYSQL_DSN")
	if dsn != "" {
		return abrirMySQL(t, dsn)
	}

	return abrirSQLite(t)
}

func abrirSQLite(t *testing.T) *gorm.DB {
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

func abrirMySQL(t *testing.T, dsn string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("abrir banco: %v", err)
	}

	err = db.AutoMigrate(
		&importacao.Escola{},
		&importacao.Endereco{},
		&importacao.Escolaridade{},
		&importacao.Turno{},
		&importacao.EscolaridadeEscola{},
		&importacao.EscolaTurno{},
		&importacao.VagasEscolares{},
	)

	return db
}
