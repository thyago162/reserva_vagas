package candidato

import (
	"errors"
	"net/http"
	"reserva_vagas/internalerrors"

	"github.com/go-chi/chi"
	"gorm.io/gorm"
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
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, http.StatusNotFound, internalerrors.ErrNotFound
	}

	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return candidato.Cpf, http.StatusOK, nil
}
