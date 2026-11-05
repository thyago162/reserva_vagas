package endpoints

import (
	"reserva_vagas/internal/escolaridades"
	"reserva_vagas/internal/escolas"
	"reserva_vagas/internal/infrastructure/db/repositories"
	"reserva_vagas/internal/login"
	"reserva_vagas/internal/reservas"
	"reserva_vagas/internal/vagasescolares"

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

	// ====================================================================
	// Módulo: Reservas Escolares
	//===========================================================================
	candRepo := repositories.NewCandidatoRepository(database)
	endRepo := repositories.NewEnderecoRepositorydb(database)
	resRepo := repositories.NewReservaRepository(database)

	reservas.RegisterRoutes(r, database, candRepo, endRepo, resRepo)

	// ========================================================================
	// Módulo: Login
	// ========================================================================
	userRepo := repositories.NewUsuarioRepository(database)
	login.RegisterRoutes(r, userRepo)

}
