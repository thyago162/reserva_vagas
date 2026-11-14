package reservas

import (
	"errors"
	"math/rand"
	"reserva_vagas/internal/candidato"
	"reserva_vagas/internal/enderecos"
	"reserva_vagas/internal/usuario"
	"reserva_vagas/internal/vagasescolares"
	"strings"

	"gorm.io/gorm"
)

type Publisher interface {
	PublicarEmailReservaAsync(email string, nome string, senha string)
}

type RealizarReservaService struct {
	Db       *gorm.DB
	ResRepo  Repository
	CandRepo candidato.Repository
	EndRepo  enderecos.Repository
	UserRepo usuario.Repository
	RabitMq  Publisher
	vagaRepo vagasescolares.Repository
}

func NewRealizarReservaService(
	db *gorm.DB,
	resRepo Repository,
	candRepo candidato.Repository,
	endRepo enderecos.Repository,
	userRepo usuario.Repository,
	rabitMq Publisher,
	vagaRepo vagasescolares.Repository,
) *RealizarReservaService {
	return &RealizarReservaService{
		Db:       db,
		ResRepo:  resRepo,
		CandRepo: candRepo,
		EndRepo:  endRepo,
		UserRepo: userRepo,
		RabitMq:  rabitMq,
		vagaRepo: vagaRepo,
	}
}

type CriarReservaDTO struct {
	EscolaId                string  `json:"escola_id"`
	EscolaridadeId          string  `json:"escolaridade_id"`
	TurnoId                 string  `json:"turno_id"`
	VagaId                  string  `json:"vaga_id"`
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
	Numero                  *string `json:"numero"`
	Complemento             string  `json:"complemento"`
	Bairro                  string  `json:"bairro"`
	Cidade                  string  `json:"cidade"`
	Uf                      string  `json:"uf"`
	Cep                     string  `json:"cep"`
	Latitude                float64 `json:"latitude"`
	Longitude               float64 `json:"longitude"`
}

type ReservaResponse struct {
	NomeCandidato           string
	CpfCandidato            string
	DtNascimentoCandidato   string
	NomeResponsavel         string
	CpfResponsavel          string
	DtNascimentoResponsavel string
	Telefone                string
	EmailResponsavel        string
	Logradouro              string
	Numero                  string
	Complemento             string
	Bairro                  string
	Cidade                  string
	Uf                      string
	Situacao                string
}

func (r *RealizarReservaService) RealizarReserva(dto CriarReservaDTO) (ReservaResponse, error) {
	var reservaResponse ReservaResponse

	err := r.Db.Transaction(func(tx *gorm.DB) error {
		err := debitarVaga(tx, r.vagaRepo, dto.VagaId)
		if err != nil {
			return errors.New("Falha ao debitar a vaga")
		}

		// 1. Salva os dados gerais do candidato/responsavel
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

		candidatoObj, err := salvarCandidato(candDto, r.CandRepo, tx)
		if err != nil {
			return err // Retorna apenas o erro
		}

		// 2. Salva o endereço
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
			EndereableID:   candidatoObj.Id,
			EndereableType: "candidato",
		}

		err = salvarEndereco(endDto, r.EndRepo, tx)
		if err != nil {
			return err // Retorna apenas o erro
		}

		reserva, err := salvarReserva(dto.EscolaId, dto.EscolaridadeId, dto.TurnoId, dto.VagaId, candidatoObj.Id, r.ResRepo, tx)
		if err != nil {
			return errors.New("Falha ao salvar a reserva")
		}

		senhaAleatoria := gerarStringAleatoria(10)
		cadastraResponsavel(candidatoObj.EmailResponavel, senhaAleatoria, r.UserRepo)

		//Joga na fila do rabitmq
		r.RabitMq.PublicarEmailReservaAsync(candidatoObj.EmailResponavel, candidatoObj.NomeResponsavel, senhaAleatoria)

		// Corrigida a formatação de datas do Go ("02/01/2006")
		reservaResponse = ReservaResponse{
			NomeCandidato:           candidatoObj.Nome,
			CpfCandidato:            candidatoObj.Cpf,
			DtNascimentoCandidato:   candidatoObj.DtNascimento.Format("02/01/2006"),
			NomeResponsavel:         candidatoObj.NomeResponsavel,
			CpfResponsavel:          candidatoObj.CpfResponsavel,
			DtNascimentoResponsavel: candidatoObj.DtNascimentoResponsavel.Format("02/01/2006"),
			Telefone:                candidatoObj.Telefone,
			EmailResponsavel:        candidatoObj.EmailResponavel,
			Logradouro:              dto.Logradouro,
			Numero:                  *dto.Numero,
			Complemento:             dto.Complemento,
			Bairro:                  dto.Bairro,
			Cidade:                  dto.Cidade,
			Uf:                      dto.Uf,
			Situacao:                string(reserva.Situacao),
		}

		return nil // Transação bem sucedida
	})

	if err != nil {
		return ReservaResponse{}, err
	}

	return reservaResponse, nil
}
func salvarCandidato(
	candDto candidato.CreateCandidatoDTO,
	candRepo candidato.Repository,
	tx *gorm.DB,
) (*candidato.Candidato, error) {
	candidato, err := candidato.NewCandidato(candDto)
	if err != nil {
		return nil, err
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
		return err
	}

	err = endRepo.StoreWithTx(tx, endereco)
	if err != nil {
		return err
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
) (*Reserva, error) {
	reserva, err := NewReserva(escolaId, EscolaridadeId, turnoId, candidatoId, vagaId)
	if err != nil {
		return nil, errors.New("Falha na validação dos dados da reserva.")
	}

	err = resRepo.StoreWithTx(tx, reserva)
	if err != nil {
		return nil, err
	}

	return reserva, nil
}

func cadastraResponsavel(emailResp, senhaAleatoria string, userRepo usuario.Repository) error {
	usuario, err := usuario.NewUsuario(emailResp, senhaAleatoria, usuario.User)
	if err != nil {
		return errors.New("Falha na validação dos dados do usuário")
	}

	err = userRepo.Store(usuario)
	if err != nil {
		return err
	}

	return nil
}

func gerarStringAleatoria(n int) string {
	const letras = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var sb strings.Builder
	for i := 0; i < n; i++ {
		idx := rand.Intn(len(letras))
		sb.WriteByte(letras[idx])
	}
	return sb.String()
}

func debitarVaga(tx *gorm.DB, vagaRepo vagasescolares.Repository, vagaId string) error {
	return vagaRepo.DebitarVaga(tx, vagaId)
}
