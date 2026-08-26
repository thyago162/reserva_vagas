package escolas

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

	if err := s.repo.Salvar(escola); err != nil {
		return err
	}

	return nil
}

func (s *Service) GetEscolas() ([]*Escola, error) {
	escolas, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	return escolas, nil
}

func (s Service) GetEscolaById(id string) (*Escola, error) {
	escola, err := s.repo.GetById(id)
	if err != nil {
		return nil, err
	}

	return escola, nil
}
