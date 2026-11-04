package reservas

import (
	"encoding/json"
	"net/http"
)

type RealizarReservaHandler struct {
	service *RealizarReservaService
}

func NewRealizarReservaHandler(s *RealizarReservaService) *RealizarReservaHandler {
	return &RealizarReservaHandler{service: s}
}

func (h *RealizarReservaHandler) RealizarReservaCandidato(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	var req CriarReservaDTO
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, http.StatusBadRequest, err
	}

	err := h.service.RealizarReserva(req)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)

	return map[string]string{"message": "Reserva realizada com sucesso."}, http.StatusCreated, nil
}
