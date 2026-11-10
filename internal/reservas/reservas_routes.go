package reservas

import (
	"reserva_vagas/internal/auth"
	"reserva_vagas/internal/candidato"
	"reserva_vagas/internal/enderecos"
	"reserva_vagas/internal/infrastructure/http"
	"reserva_vagas/internal/infrastructure/messasing"
	"reserva_vagas/internal/usuario"

	"github.com/go-chi/chi"
	"gorm.io/gorm"
)

func RegisterRoutes(
	r chi.Router,
	db *gorm.DB,
	candRepo candidato.Repository,
	endRepo enderecos.Repository,
	resRepo Repository,
	userRepo usuario.Repository,
	rabbitMq *messasing.RabbitPublisher,
) {

	service := NewRealizarReservaService(
		db,
		resRepo,
		candRepo,
		endRepo,
		userRepo,
		rabbitMq,
	)
	handler := NewRealizarReservaHandler(service)

	r.Route("/reserva-vagas", func(r chi.Router) {
		r.Use(auth.AuthMiddleware)
		r.Use(auth.RequireRole(usuario.Mcp))

		r.Post("/realizar-reserva", http.HandlerErro(handler.RealizarReservaCandidato))
	})

}
