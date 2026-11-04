package reservas

import (
	"reserva_vagas/internal/candidato"
	"reserva_vagas/internal/enderecos"
	"reserva_vagas/internal/infrastructure/http"

	"github.com/go-chi/chi"
	"gorm.io/gorm"
)

func RegisterRoutes(r chi.Router, db *gorm.DB, candRepo candidato.Repository, endRepo enderecos.Repository, resRepo Repository) {

	service := NewRealizarReservaService(db, resRepo, candRepo, endRepo)
	handler := NewRealizarReservaHandler(service)

	r.Post("/realizar-reserva", http.HandlerErro(handler.RealizarReservaCandidato))
}
