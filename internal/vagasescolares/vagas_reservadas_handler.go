package vagasescolares

import (
	"net/http"
	"reserva_vagas/internalerrors"
	"strconv"

	"github.com/go-chi/chi"
)

type VagasEscolaresHandler struct {
	service *Service
}

func NewVagasEscolaresHandler(s *Service) *VagasEscolaresHandler {
	return &VagasEscolaresHandler{service: s}
}

func (h *VagasEscolaresHandler) Index(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	id := chi.URLParam(r, "id")
	anoLetivoStr := chi.URLParam(r, "anoLetivo")

	if id == "" || anoLetivoStr == "" {
		return nil, http.StatusBadRequest, internalerrors.ErrValidation
	}

	anoLetivo, err := strconv.Atoi(anoLetivoStr)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	vagasEscolares, err := h.service.Index(id, anoLetivo)

	return vagasEscolares, http.StatusOK, nil
}
