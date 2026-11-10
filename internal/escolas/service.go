package escolas

import (
	"fmt"
	"math"
	"reserva_vagas/internal/enderecos"
	"reserva_vagas/internal/escolaridades"
	"reserva_vagas/internal/importacao"
	"reserva_vagas/internalerrors"
	"slices"
	"time"
)

const raioTerraKm = 6371.0

type Service struct {
	repo  Repository
	repoE escolaridades.Repository
}

func NewService(r Repository, repoE escolaridades.Repository) *Service {
	return &Service{repo: r, repoE: repoE}
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

func (s Service) GetEscolasByDtNascimento(dtNascimento string, latitude, longitude float64) ([]*Escola, error) {
	layoutBR := "02-01-2006"
	fmt.Println(latitude)
	fmt.Println(longitude)
	dt, err := time.Parse(layoutBR, dtNascimento)
	if err != nil {
		return nil, err
	}

	idade := calcularIdade(dt)
	infoEscolaridade, err := escolaridades.ObterEscolaridadePorIdade(idade)
	if err != nil {
		return nil, err
	}

	escolaridade, err := s.repoE.FindByCodigo(infoEscolaridade.Codigo)
	if err != nil {
		return nil, err
	}

	escolas, err := s.repo.GetByEscolaByEscolaridadeId(escolaridade.Id)
	if err != nil {
		return nil, err
	}

	escolasProx := calcularDistancia(escolas, latitude, longitude)
	return escolasProx, nil
}

func calcularIdade(dataNascimento time.Time) int {
	hoje := time.Now()
	idade := hoje.Year() - dataNascimento.Year()

	// Se ainda não chegou o mês do aniversário, ou se é o mesmo mês mas não chegou o dia: subtrai 1 ano
	if hoje.Month() < dataNascimento.Month() ||
		(hoje.Month() == dataNascimento.Month() && hoje.Day() < dataNascimento.Day()) {
		idade--
	}

	return idade
}

func calcularDistancia(escolas []*Escola, latitude, longitude float64) []*Escola {
	var escolasResult []*Escola

	for _, escola := range escolas {
		if escola.Endereco == nil {
			continue
		}

		dLat := (escola.Endereco.Latitude - latitude) * math.Pi / 180.0
		dLon := (escola.Endereco.Longitude - longitude) * math.Pi / 180.0

		radLat1 := latitude * math.Pi / 180.0
		radLat2 := escola.Endereco.Latitude * math.Pi / 180.0

		a := math.Sin(dLat/2)*math.Sin(dLat/2) +
			math.Sin(dLon/2)*math.Sin(dLon/2)*math.Cos(radLat1)*math.Cos(radLat2)

		c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
		distancia := raioTerraKm * c
		fmt.Println(distancia)

		escola.Distancia = distancia
		if distancia < 3 {
			escolasResult = append(escolasResult, escola)
		}
	}

	fmt.Println(escolasResult)

	slices.SortFunc(escolasResult, func(a, b *Escola) int {
		if a.Distancia < b.Distancia {
			return -1 // 'a' vem antes de 'b'
		}
		if a.Distancia > b.Distancia {
			return 1 // 'a' vem depois de 'b'
		}
		return 0 // distâncias iguais
	})

	return escolasResult
}
