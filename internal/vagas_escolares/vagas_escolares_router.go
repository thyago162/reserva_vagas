package vagasescolares

import (
	"reserva_vagas/internal/infrastructure/http"

	"github.com/go-chi/chi"
)

func RegisterRoutes(r chi.Router, repo Repository) {
	service := NewVagasEscolaresService(repo)
	handler := NewVagasEscolaresHandler(service)

	r.Route("/vagas-escolares", func(r chi.Router) {
		r.Get("/", http.HandlerErro(handler.Index))
	})
}
