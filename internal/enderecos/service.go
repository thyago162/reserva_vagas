package enderecos

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{repo: r}
}

func (s Service) SalvarEndereco(input CreateEnderecoInput) error {
	endereco, err := NewEndereco(input)
	if err != nil {
		return err
	}

	if err := s.repo.Store(endereco); err != nil {
		return err
	}

	return nil
}
