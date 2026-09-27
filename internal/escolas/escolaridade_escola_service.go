package escolas

import (
	"errors"
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

func (s *EscolaridadeEscolaService) AddEscolaridadeEscola(escolaridadeIds []string, escolaId string) error {
	qntIds := len(escolaridadeIds)

	if _, err := s.repo.Show(escolaId); err != nil {
		return err
	}

	qntQ, err := s.repoE.FindByIds(escolaridadeIds)

	if err != nil {
		return err
	}

	if qntQ < int64(qntIds) {
		return errors.New("uma ou mais escolaridades informadas não existem")
	}

	escolEscolas := make([]*EscolaridadeEscola, 0, qntIds)
	for _, v := range escolaridadeIds {
		escolEscola, err := NewEscolaridadeEscola(escolaId, v)
		if err != nil {
			return err
		}

		escolEscolas = append(escolEscolas, escolEscola)
	}

	return s.repoEE.Store(escolEscolas)
}

func (s *EscolaridadeEscolaService) ExibirEscolaridadesEscola(escolaId string) ([]*EscolaridadeEscola, error) {
	escolaridades, err := s.repoEE.Index(escolaId)
	if err != nil {
		return nil, err
	}

	return escolaridades, nil
}

func (s *EscolaridadeEscolaService) Remover(ids []string) error {
	err := s.repoEE.Destroy(ids)
	if err != nil {
		return err
	}

	return nil
}
