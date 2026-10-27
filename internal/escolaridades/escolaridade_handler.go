package escolaridades

import "net/http"

type EscolaridadeHandler struct {
	service *Service
}

func NewEscolaridadeHandler(s *Service) *EscolaridadeHandler {
	return &EscolaridadeHandler{service: s}
}

func (h *EscolaridadeHandler) Index(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	escolaridadesDb, err := h.service.repo.Index()
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	if escolaridadesDb == nil {
		escolaridadesDb = make([]*Escolaridade, 0)
	}

	return escolaridadesDb, http.StatusOK, nil
}
