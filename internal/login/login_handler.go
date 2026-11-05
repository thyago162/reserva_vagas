package login

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reserva_vagas/internal/auth"
)

type LoginHandler struct {
	loginService *LoginService
}

func NewLoginHandler(ls *LoginService) *LoginHandler {
	return &LoginHandler{loginService: ls}
}

type LoginRequest struct {
	Email string `json:"email"`
	Senha string `json:"senha"`
}

func (h *LoginHandler) HandleLogin(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	if r.Method != http.MethodPost {
		return nil, http.StatusMethodNotAllowed, errors.New("Método não permitido")
	}

	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		fmt.Printf("[ERRO DECODE]: %v\n", err) // Isto vai mostrar se o erro é io.EOF (corpo vazio)
		return nil, http.StatusBadRequest, err
	}

	u, err := h.loginService.Login(req.Email, req.Senha)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized) // 401 Unauthorized

		return nil, http.StatusUnauthorized, err
	}

	token, err := auth.GerarToken(u)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	w.Header().Set("Content-Type", "application/json")
	return map[string]string{
		"access_token": token,
	}, http.StatusOK, nil
}
