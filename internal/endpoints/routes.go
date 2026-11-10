package endpoints

import (
	"reserva_vagas/internal/candidato"
	"reserva_vagas/internal/escolaridades"
	"reserva_vagas/internal/escolas"
	"reserva_vagas/internal/infrastructure/db/repositories"
	"reserva_vagas/internal/infrastructure/messasing"
	"reserva_vagas/internal/login"
	"reserva_vagas/internal/reservas"
	"reserva_vagas/internal/vagasescolares"

	"github.com/go-chi/chi"
	"gorm.io/gorm"
)

func InitRoutes(r chi.Router, database *gorm.DB, rabbitConn *messasing.RabbitPublisher) {
	// ===================================================
	// INICIALIZAÇÃO DE REPOSITORIOS
	// =======================================================
	repoEscola := repositories.NewEscolaRepositoryDb(database)
	repoEE := repositories.NewEscolaridadeEscolaRepositoryDb(database)
	repoE := repositories.NewEscolaridadeRepositoryDb(database)
	repoEscolaridade := repositories.NewEscolaridadeRepositoryDb(database)
	repo := repositories.NewVagasEscolaresRepository(database)
	candRepo := repositories.NewCandidatoRepository(database)
	endRepo := repositories.NewEnderecoRepositorydb(database)
	resRepo := repositories.NewReservaRepository(database)
	userRepo := repositories.NewUsuarioRepository(database)

	// ==========================================
	// Módulo: Escolas
	// ==========================================
	escolas.RegisterRoutes(r, repoEscola, repoEE, repoE)

	// ==========================================
	// Módulo: Escolaridades
	// ==========================================
	escolaridades.RegisterRoutes(r, repoEscolaridade)

	// ========================================================
	// Módul: Vagas Escolares
	// =========================================================
	vagasescolares.RegisterRoutes(r, repo)

	// ====================================================================
	// Módulo: Reservas Escolares
	//===========================================================================
	reservas.RegisterRoutes(r, database, candRepo, endRepo, resRepo, userRepo, rabbitConn)

	// ========================================================================
	// Módulo: Login
	// ========================================================================

	login.RegisterRoutes(r, userRepo)

	// ========================================================================
	// Modulo: Candidato
	// ============================================================================

	candidato.RegisterRoutes(r, candRepo)

}
