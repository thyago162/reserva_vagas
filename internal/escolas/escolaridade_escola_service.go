package escolas

import (
	"fmt"
	"reserva_vagas/internal/escolaridades"
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
	if _, err := s.repo.Show(escolaId); err != nil {
		return err
	}

	fmt.Println("chegou 1")

	if _, err := s.repoE.Show(escolaridadeId); err != nil {
		return err
	}

	fmt.Println("chegou 2")

	escolEscola, err := NewEscolaridadeEscola(escolaId, escolaridadeId)
	if err != nil {
		return err
	}

	fmt.Println("chegou 3")

	fmt.Println(escolEscola)

	return s.repoEE.Store(escolEscola)
}

func (s *EscolaridadeEscolaService) ExibirEscolaridadesEscola(escolaId string) ([]*EscolaridadeEscola, error) {
	escolaridades, err := s.repoEE.Index(escolaId)
	if err != nil {
		return nil, err
	}

	return escolaridades, nil
}
