package escolas

import (
	"reserva_vagas/internal/endpoints"

	"github.com/go-chi/chi"
)

// O pacote escolas recebe os Repositórios/Interfaces que ele precisa, sem importar o pacote repositories!
func RegisterRoutes(r chi.Router, repo Repository, repoEE EscolaridadeEscolaRepository) {
	service := NewService(repo)
	eeService := NewEscolaridadeEscolaService(repoEE, repo, nil)

	handler := NewEscolaHandler(service)
	eeHandler := NewEscolaridadeEscolaHandler(eeService)

	r.Route("/escolas", func(r chi.Router) {
		r.Post("/", endpoints.HandlerErro(handler.Criar))
		r.Get("/", endpoints.HandlerErro(handler.Index))
		r.Get("/{id}", endpoints.HandlerErro(handler.Show))
		r.Put("/{id}", endpoints.HandlerErro(handler.Update))
		r.Delete("/{id}", endpoints.HandlerErro(handler.Destroy))

	})

	r.Route("/escolaridade-escola", func(r chi.Router) {
		r.Post("/", endpoints.HandlerErro(eeHandler.Store))
		r.Get("/", endpoints.HandlerErro(eeHandler.Index))
	})
}
