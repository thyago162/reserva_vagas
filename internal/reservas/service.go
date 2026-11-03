package reservas

import (
	"errors"
	"reserva_vagas/internalerrors"

	"gorm.io/gorm"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Index() ([]*Reserva, error) {
	var reservas []*Reserva

	reservas, err := s.repo.Index()
	if err != nil {
		return nil, err
	}

	return reservas, nil
}

func (s *Service) Store(escolaId, escolaridadeId, turnoId, vagaId, candidatoId string) error {
	reserva, err := NewReserva(escolaId, escolaridadeId, turnoId, vagaId, candidatoId)
	if err == nil {
		return internalerrors.ErrInternal
	}

	return s.repo.Store(reserva)

}

func (s *Service) Show(id string) (*Reserva, error) {
	if id == "" {
		return nil, errors.New("id não pode ser vazio")
	}

	reserva, err := s.repo.Show(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("reserva não encontrada")
		}

		return nil, err
	}

	return reserva, nil
}

func (s *Service) Destroy(id string) error {
	if id == "" {
		return errors.New("id não pode ser vazio")
	}

	_, err := s.repo.Show(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("reserva não encontrada")
		}

		return err
	}

	return s.repo.Destroy(id)
}
