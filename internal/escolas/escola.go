package escolas

import (
	"reserva_vagas/internal/enderecos"
	"reserva_vagas/internal/vagasescolares"
	"reserva_vagas/internalerrors"
	"time"

	"github.com/google/uuid"
)

type Escola struct {
	Id                 string                          `gorm:"primaryKey;size:36;not null"`
	Nome               string                          `validate:"min=5,max=150,required" gorm:"not null"`
	Inep               string                          `validate:"min=6,required" gorm:"not null"`
	Email              string                          `validate:"email,required" gorm:"not null"`
	Sigla              string                          `validate:"max=10"`
	Telefone           string                          `validate:"required,max=15" gorm:"not null"`
	CreatedAt          time.Time                       `gorm:"autoCreateTime"`
	UpdatedAt          time.Time                       `gorm:"autoUpdateTime"`
	EscolaridadeEscola []EscolaridadeEscola            `gorm:"foreignKey:EscolaId"`
	EscolaTurno        []EscolaTurno                   `gorm:"foreignKey:EscolaId"`
	Endereco           *enderecos.Endereco             `gorm:"polymorphic:Endereable;polymorphicValue:escola"`
	VagasEscolares     []vagasescolares.VagasEscolares `gorm:"foreignKey:EscolaId"`
	Distancia          float64                         `json:"distancia" gorm:"-"`
}

type EscolaResumoDB struct {
	EscolaId         string `gorm:"column:escola_id"`
	NomeEscola       string `gorm:"column:escola_nome"`
	Inep             string `gorm:"column:inep"`
	TelefoneEscola   string `gorm:"column:telefone"`
	EmailEscola      string `gorm:"email_escola"`
	EnderecoId       string `gorm:"column:endereco_id"`
	TipoLogradouro   string `gorm:"column:tipo_logradouro"`
	Logradouro       string `gorm:"column:logradouro"`
	Numero           string `gorm:"column:numero"`
	Complemento      string `gorm:"column:complemento"`
	Bairro           string `gorm:"column:bairro"`
	Cidade           string `gorm:"column:cidade"`
	Uf               string `gorm:"column:uf"`
	Cep              string `gorm:"column:cep"`
	IdEscolaridade   string `gorm:"column:escolaridade_id"`
	NomeEscolaridade string `gorm:"column:escolaridade_nome"`
	IdTurno          string `gorm:"column:turno_id"`
	NomeTurno        string `gorm:"column:nome_turno"`
	VagasDisponiveis int32  `gorm:"column:vagas_disponiveis"`
}

func NewEscola(nome, inep, email, sigla, telefone string) (*Escola, error) {
	escola := &Escola{
		Id:       uuid.NewString(),
		Nome:     nome,
		Inep:     inep,
		Email:    email,
		Sigla:    sigla,
		Telefone: telefone,
	}

	err := internalerrors.ValidateStruct(escola)
	if err == nil {
		return escola, nil
	}

	return nil, err
}
