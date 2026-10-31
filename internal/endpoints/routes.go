package endpoints

import (
	"reserva_vagas/internal/escolaridades"
	"reserva_vagas/internal/escolas"
	"reserva_vagas/internal/infrastructure/db/repositories"
	vagasescolares "reserva_vagas/internal/vagas_escolares"

	"github.com/go-chi/chi"
	"gorm.io/gorm"
)

func InitRoutes(r chi.Router, database *gorm.DB) {
	// ==========================================
	// Módulo: Escolas
	// ==========================================
	repoEscola := repositories.NewEscolaRepositoryDb(database)
	repoEE := repositories.NewEscolaridadeEscolaRepositoryDb(database)
	repoE := repositories.NewEscolaridadeRepositoryDb(database)

	escolas.RegisterRoutes(r, repoEscola, repoEE, repoE)

	// ==========================================
	// Módulo: Escolaridades
	// ==========================================
	repoEscolaridade := repositories.NewEscolaridadeRepositoryDb(database)
	escolaridades.RegisterRoutes(r, repoEscolaridade)

	// ========================================================
	// Módul: Vagas Escolares
	// =========================================================
	repo := repositories.NewVagasEscolaresRepository(database)
	vagasescolares.RegisterRoutes(r, repo)
}
