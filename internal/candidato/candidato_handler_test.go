package candidato_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"reserva_vagas/internal/candidato"
	"reserva_vagas/internal/infrastructure/db/repositories"
	"reserva_vagas/internal/testutil"
	"testing"

	"github.com/go-chi/chi"
)

func TestCandidatoHandler_FindByCpf_RetornaReservaExiste(t *testing.T) {
	db := testutil.NovoDBTeste(t)
	repo := repositories.NewCandidatoRepository(db)
	service := candidato.NewService(repo)
	h := candidato.NewCandidatoHandler(service)

	req := httptest.NewRequest(http.MethodGet, "/candidatos/52998224725", nil)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("cpf", "52998224725")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	rec := httptest.NewRecorder()
	body, status, err := h.FindByCpf(rec, req)

	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}

	if status != http.StatusOK {
		t.Fatalf("status = %d, esperado %d", status, http.StatusOK)
	}

	existe, ok := body.(bool)
	if !ok {
		t.Fatalf("resposta deveria ser só o 1, veio %T: %+v", body, body)
	}

	if !existe {
		t.Errorf("existe = %t, esperado True", existe)
	}
}
