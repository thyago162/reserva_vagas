package escolaridades

import (
	"reserva_vagas/internal/infrastructure/http"

	"github.com/go-chi/chi"
)

func RegisterRoutes(r chi.Router, repo Repository) {
	service := NewService(repo)
	handler := NewEscolaridadeHandler(service)

	r.Route("/escolaridades", func(r chi.Router) {
		r.Get("/", http.HandlerErro(handler.Index))
	})
}
