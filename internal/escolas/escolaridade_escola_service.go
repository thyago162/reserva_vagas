package escolas

import (
	"reserva_vagas/internal/escolaridades"
	"reserva_vagas/internalerrors"
)

type EscolaridadeEscolaService struct {
	repoEE EscolaridadeEscolaRepository
	repo   Repository
	repoE  escolaridades.Repository
}

func NewEscolaridadeEscolaService(ree EscolaridadeEscolaRepository, r Repository, re escolaridades.Repository) *EscolaridadeEscolaService {
	return &EscolaridadeEscolaService{repoEE: ree, repo: r, repoE: re}
}

func (s *EscolaridadeEscolaService) AddEscolaridadeEscola(escolaridadeId, escolaId string) error {
	escola, err := s.repo.Show(escolaId)

	if err != nil {
		return err
	}

	escolaridade, err := s.repoE.Show(escolaridadeId)
	if err != nil {
		return err
	}

	exists, err := s.repoEE.Exists(escolaridadeId, escolaId)
	if err != nil {
		return err
	}

	if exists {
		return internalerrors.ErrRegistroExiste
	}

	escolaridadeEscola, err := NewEscolaridadeEscola(escola.Id, escolaridade.Id)
	if err != nil {
		return err
	}

	if err := s.repoEE.Store(escolaridadeEscola); err != nil {
		return err
	}

	return nil
}

func (s *EscolaridadeEscolaService) ExibirEscolaridadesEscola(escolaId string) ([]*EscolaridadeEscola, error) {
	escolaridades, err := s.repoEE.Index(escolaId)
	if err != nil {
		return nil, err
	}

	return escolaridades, nil
}
