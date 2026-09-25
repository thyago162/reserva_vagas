package escolas

import (
	"encoding/json"
	"errors"
	"net/http"
	"reserva_vagas/internalerrors"
)

type EscolaridadeEscolaHandler struct {
	service *EscolaridadeEscolaService
}

func NewEscolaridadeEscolaHandler(s *EscolaridadeEscolaService) *EscolaridadeEscolaHandler {
	return &EscolaridadeEscolaHandler{service: s}
}

type CriarUpdateEscolEscolaRequest struct {
	EscolaID       string `json:"escola_id"`
	EscolaridadeID string `json:"escolaridade_id"`
}

func (h *EscolaridadeEscolaHandler) Store(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	var req CriarUpdateEscolEscolaRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, http.StatusBadRequest, err
	}

	err := h.service.AddEscolaridadeEscola(req.EscolaridadeID, req.EscolaID)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	return map[string]string{"message": "Escolaridade cadastrada com sucesso"}, http.StatusCreated, nil
}

func (h *EscolaridadeEscolaHandler) Index(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	escolaID := r.URL.Query().Get("escolaId")
	if escolaID == "" {
		return nil, http.StatusBadRequest, errors.New("o parâmetro escolaId é obrigatório")
	}

	escolaridadesEscola, err := h.service.ExibirEscolaridadesEscola(escolaID)
	if err != nil {
		if errors.Is(err, internalerrors.ErrRegistroExiste) {
			return nil, http.StatusConflict, err // Status 409 para registro duplicado
		}
		return nil, http.StatusBadRequest, err
	}

	if escolaridadesEscola == nil {
		escolaridadesEscola = make([]*EscolaridadeEscola, 0)
	}

	return escolaridadesEscola, http.StatusOK, nil
}
