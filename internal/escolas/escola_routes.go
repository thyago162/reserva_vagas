package escolas

import (
	"reserva_vagas/internal/auth"
	"reserva_vagas/internal/escolaridades"
	"reserva_vagas/internal/infrastructure/http"
	"reserva_vagas/internal/usuario"

	"github.com/go-chi/chi"
)

func RegisterRoutes(r chi.Router, repo Repository, repoEE EscolaridadeEscolaRepository, repoE escolaridades.Repository) {
	service := NewService(repo, repoE)
	eeService := NewEscolaridadeEscolaService(repoEE, repo, repoE)

	handler := NewEscolaHandler(service)
	eeHandler := NewEscolaridadeEscolaHandler(eeService)

	r.Route("/escolas", func(r chi.Router) {

		// =================================================================================
		//  Rotas MCP Server
		// =================================================================================
		r.Use(auth.AuthMiddleware)
		r.Use(auth.RequireRole(usuario.Mcp))

		r.Get("/get-escolas-dt-nascimento", http.HandlerErro(handler.GetEscolasByDtNascimento))

		// ===========================================================================
		// Rotas protegidas
		// ============================================================================

		r.Group(func(r chi.Router) {
			// =================================================================================
			// Rotas da Escola
			// ==================================================================================

			// Validação do token JWT
			r.Use(auth.AuthMiddleware)

			// Rotas do Admin e Escola
			r.Use(auth.RequireRole(usuario.Admin, usuario.Escola))

			// ======================================== ROTAS ========================================

			r.Get("/{id}", http.HandlerErro(handler.Show))
			r.Put("/{id}", http.HandlerErro(handler.Update))
		})

		r.Group(func(r chi.Router) {
			// =================================================================================
			// Rotas do ADMIN
			// ==================================================================================

			// Validação do token JWT
			r.Use(auth.AuthMiddleware)
			r.Use(auth.RequireRole(usuario.Admin))

			r.Get("/", http.HandlerErro(handler.Index))
			r.Post("/", http.HandlerErro(handler.Criar))
			r.Post("/importar", http.HandlerErro(handler.ImportarEmLotes))
			r.Delete("/{id}", http.HandlerErro(handler.Destroy))
		})
	})

	r.Route("/escolaridade-escola", func(r chi.Router) {
		// =================================================================================
		// Rotas da Escola
		// ==================================================================================

		// Validação do token JWT
		r.Use(auth.AuthMiddleware)

		// Rotas do Admin e Escola
		r.Use(auth.RequireRole(usuario.Admin, usuario.Escola))

		r.Post("/", http.HandlerErro(eeHandler.Store))
		r.Get("/", http.HandlerErro(eeHandler.Index))
	})
}
