package escolaridades

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{repo: r}
}

func (s *Service) GetEscolas() ([]*Escolaridade, error) {
	escolas, err := s.repo.Index()
	if err != nil {
		return nil, err
	}

	return escolas, nil
}
