package escolas

import "reserva_vagas/internalerrors"

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{repo: r}
}

func (s *Service) SalvarEscola(nome string, inep string, email string, sigla string) error {
	escola, err := NewEscola(nome, inep, email, sigla)
	if err != nil {
		return err
	}

	if err := s.repo.Store(escola); err != nil {
		return err
	}

	return nil
}

func (s *Service) GetEscolas() ([]*Escola, error) {
	escolas, err := s.repo.Index()
	if err != nil {
		return nil, err
	}

	return escolas, nil
}

func (s Service) GetEscolaById(id string) (*Escola, error) {
	escola, err := s.repo.Show(id)
	if err != nil {
		return nil, err
	}

	return escola, nil
}

func (s Service) UpdateEscola(id, nome, inep, email, sigla string) error {
	escola, err := s.repo.Show(id)
	if err != nil {
		return err
	}

	escola.Nome = nome
	escola.Inep = inep
	escola.Email = email
	escola.Sigla = sigla

	if err := internalerrors.ValidateStruct(escola); err != nil {
		return err
	}

	return s.repo.Update(escola)
}

func (s Service) Destroy(id string) error {
	return s.repo.Destroy(id)
}
