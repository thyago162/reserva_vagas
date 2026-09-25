package escolas

import (
	"reserva_vagas/internal/infrastructure/http"

	"github.com/go-chi/chi"
)

func RegisterRoutes(r chi.Router, repo Repository, repoEE EscolaridadeEscolaRepository) {
	service := NewService(repo)
	eeService := NewEscolaridadeEscolaService(repoEE, repo, nil)

	handler := NewEscolaHandler(service)
	eeHandler := NewEscolaridadeEscolaHandler(eeService)

	r.Route("/escolas", func(r chi.Router) {
		r.Post("/", http.HandlerErro(handler.Criar))
		r.Get("/", http.HandlerErro(handler.Index))
		r.Get("/{id}", http.HandlerErro(handler.Show))
		r.Put("/{id}", http.HandlerErro(handler.Update))
		r.Delete("/{id}", http.HandlerErro(handler.Destroy))

	})

	r.Route("/escolaridade-escola", func(r chi.Router) {
		r.Post("/", http.HandlerErro(eeHandler.Store))
		r.Get("/", http.HandlerErro(eeHandler.Index))
	})
}
