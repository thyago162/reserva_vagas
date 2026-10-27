package escolas

import (
	"reserva_vagas/internal/enderecos"
	"reserva_vagas/internal/importacao"
	"reserva_vagas/internalerrors"
)

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{repo: r}
}

func (s *Service) SalvarEscola(nome string, inep string, email string, sigla string, telefone string, inputEndereco enderecos.CreateEnderecoInput) error {
	escola, err := NewEscola(nome, inep, email, sigla, telefone)
	if err != nil {
		return err
	}

	endereco, err := enderecos.NewEndereco(inputEndereco)
	if err != nil {
		return err
	}

	escola.Endereco = endereco

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

func (s Service) UpdateEscola(id, nome, inep, email, sigla, telefone string) error {
	escola, err := s.repo.Show(id)
	if err != nil {
		return err
	}

	escola.Nome = nome
	escola.Inep = inep
	escola.Email = email
	escola.Sigla = sigla
	escola.Telefone = telefone

	if err := internalerrors.ValidateStruct(escola); err != nil {
		return err
	}

	return s.repo.Update(escola)
}

func (s Service) Destroy(id string) error {
	return s.repo.Destroy(id)
}

func (s Service) ImportarEscolas(escolas []importacao.Escola) error {
	err := s.repo.ImportarEmLotes(escolas)
	if err != nil {
		return err
	}

	return nil
}
