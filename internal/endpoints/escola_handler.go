package endpoints

import (
	"encoding/json"
	"net/http"
	"reserva_vagas/internal/escolas"
)

type EscolaHandler struct {
	service *escolas.Service
}

func NewEscolaHandler(s *escolas.Service) *EscolaHandler {
	return &EscolaHandler{service: s}
}

type CriarEscolaRequest struct {
	Nome  string `json:"nome"`
	Inep  string `json:"inep"`
	Email string `json:"email"`
	Sigla string `json:"sigla"`
}

func (h *EscolaHandler) Criar(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	var req CriarEscolaRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, http.StatusBadRequest, err
	}

	err := h.service.SalvarEscola(req.Nome, req.Inep, req.Email, req.Sigla)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	return map[string]string{"message": "Escola cadastrada com sucesso"}, http.StatusCreated, nil
}

func (h *EscolaHandler) Index(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	escolasDb, err := h.service.GetEscolas()
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	if escolasDb == nil {
		escolasDb = make([]*escolas.Escola, 0)
	}

	return escolasDb, http.StatusOK, nil
}
