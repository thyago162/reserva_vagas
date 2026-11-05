package candidato

import (
	"reserva_vagas/internal/infrastructure/http"

	"github.com/go-chi/chi"
)

func RegisterRoutes(r chi.Router, repo Repository) {
	service := NewService(repo)
	handler := NewCandidatoHandler(service)

	r.Route("/candidato", func(r chi.Router) {
		//=============================================================
		// Rotas publicas
		//==============================================================
		r.Get("/candidato-possui-reserva/{cpf}", http.HandlerErro(handler.FindByCpf))
	})

}
