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

func (h *EscolaHandler) Criar(w http.ResponseWriter, r *http.Request) {
	var req CriarEscolaRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	err := h.service.SalvarEscola(req.Nome, req.Inep, req.Email, req.Sigla)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Escola cadastrada com sucesso"})
}

func (h *EscolaHandler) Index(w http.ResponseWriter, r *http.Request) {
	escolasDb, err := h.service.GetEscolas()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if escolasDb == nil {
		escolasDb = make([]*escolas.Escola, 0)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(escolasDb)
}
