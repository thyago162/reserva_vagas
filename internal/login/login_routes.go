package login

import (
	"reserva_vagas/internal/infrastructure/http"
	"reserva_vagas/internal/usuario"

	"github.com/go-chi/chi"
)

func RegisterRoutes(r chi.Router, repo usuario.Repository) {
	service := NewLoginService(repo)
	handler := NewLoginHandler(service)

	r.Post("/login", http.HandlerErro(handler.HandleLogin))
}
