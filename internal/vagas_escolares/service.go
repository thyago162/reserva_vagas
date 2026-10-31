package vagasescolares

import "reserva_vagas/internalerrors"

type Service struct {
	repo Repository
}

func NewVagasEscolaresService(r Repository) *Service {
	return &Service{repo: r}
}

func (s *Service) Index(escolaId string, anoLetivo int) ([]*VagasEscolares, error) {
	vagasEscolares, err := s.repo.Index(escolaId, anoLetivo)
	if err != nil {
		return nil, err
	}

	return vagasEscolares, nil
}

func (s *Service) AddVagaReservada(vagaId string) error {

	vagaEscolar, err := s.repo.Show(vagaId)
	if err != nil {
		return err
	}

	if vagaEscolar.VagasDisponiveis <= 0 {
		return internalerrors.ErrVagasIndisponiveis
	}

	vagaEscolar.VagasReservadas++
	vagaEscolar.VagasDisponiveis = vagaEscolar.VagasTotais - vagaEscolar.VagasReservadas

	return s.repo.Update(vagaEscolar)
}

func (s *Service) SubVagaReservada(vagaId string) error {

	vagaEscolar, err := s.repo.Show(vagaId)
	if err != nil {
		return err
	}

	if vagaEscolar.VagasReservadas <= 0 {
		return internalerrors.ErrVagasReservadasIndisponiveis
	}

	vagaEscolar.VagasReservadas--
	vagaEscolar.VagasDisponiveis = vagaEscolar.VagasTotais - vagaEscolar.VagasReservadas

	return s.repo.Update(vagaEscolar)
}

func (s *Service) AddVagasTotais(vagaId string, qtde int) error {

	vagaEscolar, err := s.repo.Show(vagaId)
	if err != nil {
		return err
	}

	vagaEscolar.VagasTotais += int32(qtde)
	vagaEscolar.VagasDisponiveis = vagaEscolar.VagasTotais - vagaEscolar.VagasReservadas

	return s.repo.Update(vagaEscolar)
}
