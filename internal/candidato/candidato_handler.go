package candidato

import (
	"fmt"
	"net/http"
	"reserva_vagas/internalerrors"

	"github.com/go-chi/chi"
)

type CandidatoHandler struct {
	service *Service
}

type ExisteReservaResponse struct {
	Existe bool `json:"existe"`
}

func NewCandidatoHandler(s *Service) *CandidatoHandler {
	return &CandidatoHandler{service: s}
}

func (h *CandidatoHandler) FindByCpf(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	cpf := chi.URLParam(r, "cpf")
	if cpf == "" {
		return nil, http.StatusBadRequest, internalerrors.ErrValidation
	}

	existe, err := h.service.FindByCpf(cpf)

	fmt.Println(err)

	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	fmt.Println(existe)

	return ExisteReservaResponse{Existe: existe}, http.StatusOK, nil
}
