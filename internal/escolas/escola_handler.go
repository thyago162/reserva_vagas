package escolas

import (
	"encoding/json"
	"net/http"
	"reserva_vagas/internalerrors"

	"github.com/go-chi/chi"
)

type EscolaHandler struct {
	service *Service
}

func NewEscolaHandler(s *Service) *EscolaHandler {
	return &EscolaHandler{service: s}
}

type CriarUpdateEscolaRequest struct {
	Nome  string `json:"nome"`
	Inep  string `json:"inep"`
	Email string `json:"email"`
	Sigla string `json:"sigla"`
}

func (h *EscolaHandler) Criar(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	var req CriarUpdateEscolaRequest
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
		escolasDb = make([]*Escola, 0)
	}

	return escolasDb, http.StatusOK, nil
}

func (h *EscolaHandler) Show(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	id := chi.URLParam(r, "id")
	if id == "" {
		return nil, http.StatusBadRequest, internalerrors.ErrValidation
	}

	escola, err := h.service.GetEscolaById(id)
	if err != nil {
		return nil, http.StatusNotFound, err
	}

	return escola, http.StatusOK, nil
}

func (h *EscolaHandler) Update(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	id := chi.URLParam(r, "id")
	if id == "" {
		return nil, http.StatusBadRequest, internalerrors.ErrInternal
	}

	var req CriarUpdateEscolaRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, http.StatusBadRequest, err
	}

	err := h.service.UpdateEscola(id, req.Nome, req.Inep, req.Email, req.Sigla)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	return map[string]string{"message": "Escola atualizada com sucesso"}, http.StatusOK, nil
}

func (h *EscolaHandler) Destroy(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	id := chi.URLParam(r, "id")
	if id == "" {
		return nil, http.StatusBadRequest, internalerrors.ErrValidation
	}

	err := h.service.Destroy(id)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	return map[string]string{"message": "Escola atualizada com sucesso"}, http.StatusOK, nil
}
