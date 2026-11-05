package candidato

import (
	"net/http"
	"reserva_vagas/internalerrors"

	"github.com/go-chi/chi"
)

type CandidatoHandler struct {
	service *Service
}

func NewCandidatoHandler(s *Service) *CandidatoHandler {
	return &CandidatoHandler{service: s}
}

func (h *CandidatoHandler) FindByCpf(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	cpf := chi.URLParam(r, "cpf")
	if cpf == "" {
		return nil, http.StatusBadRequest, internalerrors.ErrValidation
	}

	candidato, err := h.service.FindByCpf(cpf)
	if err != nil {
		return nil, http.StatusNotFound, err
	}

	return candidato, http.StatusOK, nil
}
