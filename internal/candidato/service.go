package candidato

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) Index() ([]*Candidato, error) {
	return s.repo.Index()
}

func (s *Service) Store(candidato *Candidato) error {
	return s.repo.Store(candidato)
}

func (s *Service) Show(id string) (*Candidato, error) {
	return s.repo.Show(id)
}

func (s *Service) Update(candidato *Candidato) error {
	return s.repo.Update(candidato)
}

func (s *Service) Delete(id string) error {
	return s.repo.Delete(id)
}
