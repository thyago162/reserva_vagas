package testutil

import (
	"fmt"
	"reserva_vagas/internal/candidato"
	"reserva_vagas/internal/importacao"
	"reserva_vagas/internal/reservas"
	"reserva_vagas/internal/usuario"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NovoDBTeste(t *testing.T) *gorm.DB {
	t.Helper()
	dsn := "root:l1nux123.@tcp(127.0.0.1:3306)/reservas_test?charset=utf8mb4&parseTime=True&loc=Local"

	return AbrirMySQL(t, dsn)
}

func AbrirMySQL(t *testing.T, dsn string) *gorm.DB {
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
		&candidato.Candidato{},
		&reservas.Reserva{},
		&usuario.Usuario{},
	)

	limparTabelas(t, db)

	return db
}

func limparTabelas(t *testing.T, db *gorm.DB) {
	t.Helper()
	tabelas := []string{"reservas", "usuarios", "candidatos", "escolaridade_escolas", "escola_turnos", "vagas_escolares", "usuarios", "enderecos", "escolas"}

	err := db.Connection(func(conn *gorm.DB) error {
		// tudo aqui usa a MESMA conexão do pool
		if err := conn.Exec("SET FOREIGN_KEY_CHECKS = 0").Error; err != nil {
			return err
		}
		for _, tb := range tabelas {
			if err := conn.Exec("TRUNCATE TABLE " + tb).Error; err != nil {
				return fmt.Errorf("truncate %s: %w", tb, err)
			}
		}
		return conn.Exec("SET FOREIGN_KEY_CHECKS = 1").Error
	})
	if err != nil {
		t.Fatalf("limpar banco: %v", err)
	}
}
