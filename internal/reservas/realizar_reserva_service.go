package reservas

import (
	"errors"
	"reserva_vagas/internal/candidato"
	"reserva_vagas/internal/enderecos"

	"gorm.io/gorm"
)

type RealizarReserva struct {
	Db       *gorm.DB
	ResRepo  Repository
	CandRepo candidato.Repository
	EndRepo  enderecos.Repository
}

func NewRealizarReserva(db *gorm.DB, resRepo Repository, candRepo candidato.Repository, endRepo enderecos.Repository) *RealizarReserva {
	return &RealizarReserva{
		Db:       db,
		ResRepo:  resRepo,
		CandRepo: candRepo,
		EndRepo:  endRepo,
	}
}

type CriarReservaDTO struct {
	EscolaId                string  `json:"escola_id"`
	EscolaridadeId          string  `json:"escolaridade_id"`
	TurnoId                 string  `json:"turno_id"`
	VagaId                  string  `json:"vaga_id"`
	CandidatoId             string  `json:"candidato_id"`
	NomeCandidato           string  `json:"nome_candidato"`
	CpfCandidato            string  `json:"cpf_candidato"`
	DtNascimentoCandidato   string  `json:"dt_nascimento_candidato"`
	NomeResponsavel         string  `json:"nome_responsavel"`
	CpfResponsavel          string  `json:"cpf_responsavel"`
	DtNascimentoResponsavel string  `json:"dt_nascimento_responsavel"`
	Telefone                string  `json:"telefone"`
	EmailResponsavel        string  `json:"email_responsavel"`
	TipoLogradouro          string  `json:"tipo_logradouro"`
	Logradouro              string  `json:"logradouro"`
	Numero                  string  `json:"numero"`
	Complemento             string  `json:"complemento"`
	Bairro                  string  `json:"bairro"`
	Cidade                  string  `json:"cidade"`
	Uf                      string  `json:"uf"`
	Cep                     string  `json:"cep"`
	Latitude                float64 `json:"latitude"`
	Longitude               float64 `json:"longitude"`
}

func (r *RealizarReserva) RealizarReserva(dto CriarReservaDTO) error {
	return r.Db.Transaction(func(tx *gorm.DB) error {

		//1. Salva os dados gerais do candidato/responsavel
		candDto := candidato.CreateCandidatoDTO{
			Nome:                    dto.NomeCandidato,
			Cpf:                     dto.CpfCandidato,
			DtNascimento:            dto.DtNascimentoCandidato,
			NomeResponsavel:         dto.NomeResponsavel,
			CpfResponsavel:          dto.CpfResponsavel,
			Telefone:                dto.Telefone,
			DtNascimentoResponsavel: dto.DtNascimentoResponsavel,
			EmailResponavel:         dto.EmailResponsavel,
		}

		candidato, err := salvarCandidato(candDto, r.CandRepo, tx)

		if err != nil {
			return err
		}

		//2. Salva o endereço
		endDto := enderecos.CreateEnderecoInput{
			TipoLogradouro: dto.TipoLogradouro,
			Logradouro:     dto.Logradouro,
			Numero:         dto.Numero,
			Complemento:    dto.Complemento,
			Bairro:         dto.Bairro,
			Cidade:         dto.Cidade,
			Uf:             dto.Uf,
			Cep:            dto.Cep,
			Latitude:       dto.Latitude,
			Longitude:      dto.Longitude,
			EndereableID:   candidato.Id,
			EndereableType: "candidato",
		}

		err = salvarEndereco(endDto, r.EndRepo, tx)
		if err != nil {
			return errors.New("Falha ao salvar o endereço.")
		}

		err = salvarReserva(dto.EscolaId, dto.EscolaridadeId, dto.TurnoId, candidato.Id, dto.VagaId, r.ResRepo, tx)
		if err != nil {
			return errors.New("Falha ao salvar a reserva")
		}

		return nil
	})
}

func salvarCandidato(
	candDto candidato.CreateCandidatoDTO,
	candRepo candidato.Repository,
	tx *gorm.DB,
) (*candidato.Candidato, error) {
	candidato, err := candidato.NewCandidato(candDto)
	if err != nil {
		return nil, errors.New("Problema na validação dos dados")
	}

	err = candRepo.StoreWithTx(tx, candidato)
	if err != nil {
		return nil, err
	}

	return candidato, err
}

func salvarEndereco(
	endDto enderecos.CreateEnderecoInput,
	endRepo enderecos.Repository,
	tx *gorm.DB,
) error {
	endereco, err := enderecos.NewEndereco(endDto)
	if err != nil {
		return errors.New("Falha na validação do endereço")
	}

	err = endRepo.StoreWithTx(tx, endereco)
	if err != nil {
		return errors.New("Falha ao salvar o endereço.")
	}

	return nil
}

func salvarReserva(
	escolaId,
	EscolaridadeId,
	turnoId,
	candidatoId,
	vagaId string,
	resRepo Repository,
	tx *gorm.DB,
) error {
	reserva, err := NewReserva(escolaId, EscolaridadeId, turnoId, candidatoId, vagaId)
	if err != nil {
		return errors.New("Falha na validação dos dados da reserva.")
	}

	err = resRepo.StoreWithTx(tx, reserva)
	if err != nil {
		return errors.New("Falha ao salvar a reserva.")
	}

	return nil
}
