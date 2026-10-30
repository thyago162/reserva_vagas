package importacao

import (
	"errors"
	"fmt"
	"reserva_vagas/internalerrors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var db *gorm.DB

type Escola struct {
	Id        string `gorm:"primaryKey;size:36"`
	Nome      string
	Inep      string
	Email     string
	Sigla     string
	Telefone  string
	Endereco  Endereco `gorm:"polymorphic:Endereable;polymorphicValue:escola"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Endereco struct {
	Id             string `gorm:"primaryKey;size:36"`
	TipoLogradouro string
	Logradouro     string
	Numero         string
	Complemento    string
	Bairro         string
	Cidade         string
	Uf             string
	Cep            string
	Latitude       float64
	Longitude      float64
	EndereableID   string `gorm:"size:36"`
	EndereableType string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type Escolaridade struct {
	Id     string `gorm:"primaryKey;size:36"`
	Codigo string
}

type Turno struct {
	Id     string `gorm:"primaryKey;size:36"`
	Codigo string
}

type EscolaridadeEscola struct {
	Id             string `gorm:"primaryKey;size:36"`
	EscolaId       string
	EscolaridadeId string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type EscolaTurno struct {
	Id        string `gorm:"primaryKey;size:36"`
	EscolaId  string
	TurnoId   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type VagasEscolares struct {
	Id             string `gorm:"primaryKey;size:36"`
	EscolaId       string
	EscolaridadeId string
	TurnoId        string
	Qtde           int32
	AnoLetivo      int32
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type VagasDTO struct {
	TurnoCodigo        string `json:"turno_codigo"`
	EscolaridadeCodigo string `json:"escolaridade_codigo"`
	Qtde               int32  `json:"qtde"`
	AnoLetivo          int32  `json:"ano_letivo"`
}

type ImportacaoEscolaDTO struct {
	Nome                 string     `json:"nome" validate:"required,max=150"`
	Inep                 string     `json:"inep" validate:"min=6,required"`
	Email                string     `json:"email" validate:"email,required"`
	Sigla                string     `json:"sigla" validate:"required,max=10"`
	Telefone             string     `json:"telefone" validate:"required,max=15"`
	TipoLogradouro       string     `json:"tipo_logradouro" validate:"required,max=50"`
	Logradouro           string     `json:"logradouro" validate:"required,max=150"`
	Numero               string     `json:"numero" validate:"required,max=10"`
	Complemento          string     `json:"complemento" validate:"max=50"`
	Bairro               string     `json:"bairro" validate:"required,max=50"`
	Cidade               string     `json:"cidade" validate:"required,max=50"`
	Uf                   string     `json:"uf" validate:"required,len=2"`
	Cep                  string     `json:"cep" validate:"required,max=9"`
	Latitude             float64    `json:"latitude" validate:"required"`
	Longitude            float64    `json:"longitude" validate:"required"`
	EscolaridadesCodigos []string   `json:"escolaridades_codigos"`
	TurnosCodigos        []string   `json:"turnos_codigos"`
	Vagas                []VagasDTO `json:"vagas"`
}

type ResultadoImportacao struct {
	Inep    string
	Nome    string
	Sucesso bool
	Erro    error
}

func SetDatabase(database *gorm.DB) {
	db = database
}

func NewImportacaoEscola(input ImportacaoEscolaDTO) (*Escola, error) {
	escolaID := uuid.New().String()
	enderecoID := uuid.New().String()
	if escolaID == "" {
		return nil, errors.New("falha ao gerar ID da escola")
	}

	importacaoEscola := &Escola{
		Id:       escolaID,
		Nome:     input.Nome,
		Inep:     input.Inep,
		Email:    input.Email,
		Sigla:    input.Sigla,
		Telefone: input.Telefone,
		Endereco: Endereco{
			Id:             enderecoID,
			EndereableID:   escolaID,
			EndereableType: "escola",
			TipoLogradouro: input.TipoLogradouro,
			Logradouro:     input.Logradouro,
			Numero:         input.Numero,
			Complemento:    input.Complemento,
			Bairro:         input.Bairro,
			Cidade:         input.Cidade,
			Uf:             input.Uf,
			Cep:            input.Cep,
			Latitude:       input.Latitude,
			Longitude:      input.Longitude,
		},
	}

	err := internalerrors.ValidateStruct(importacaoEscola)
	if err == nil {
		return importacaoEscola, nil
	}

	return nil, err
}

func NewEscolaridadeEscola(escolaId, escolaridadeId string) (*EscolaridadeEscola, error) {
	escolaridadeEscola := &EscolaridadeEscola{
		Id:             uuid.NewString(),
		EscolaId:       escolaId,
		EscolaridadeId: escolaridadeId,
	}

	err := internalerrors.ValidateStruct(escolaridadeEscola)
	if err == nil {
		return escolaridadeEscola, nil
	}

	return nil, err
}

func NewEscolaTurno(escolaId, turnoId string) (*EscolaTurno, error) {
	escolaTurno := &EscolaTurno{
		Id:       uuid.New().String(),
		EscolaId: escolaId,
		TurnoId:  turnoId,
	}

	err := internalerrors.ValidateStruct(escolaTurno)
	if err != nil {
		return nil, err
	}

	return escolaTurno, nil
}

func NewVagaEscolar(escolaId, escolaridadeId, turnoId string, qtde, anoLetivo int32) (*VagasEscolares, error) {
	vagaEscolar := &VagasEscolares{
		Id:             uuid.New().String(),
		EscolaId:       escolaId,
		EscolaridadeId: escolaridadeId,
		TurnoId:        turnoId,
		Qtde:           qtde,
		AnoLetivo:      anoLetivo,
	}

	err := internalerrors.ValidateStruct(vagaEscolar)
	if err != nil {
		return nil, err
	}

	return vagaEscolar, nil
}

func ImportarListaEscolasConcorrente(listaEscolas []ImportacaoEscolaDTO, numWorkers int) []ResultadoImportacao {
	jobs := make(chan ImportacaoEscolaDTO, len(listaEscolas))
	results := make(chan ResultadoImportacao, len(listaEscolas))

	for w := 1; w <= numWorkers; w++ {
		go workerImportacao(db, jobs, results)
	}

	for _, escolaDTO := range listaEscolas {
		jobs <- escolaDTO
	}

	close(jobs)

	var relatorio []ResultadoImportacao
	for i := 0; i < len(listaEscolas); i++ {
		relatorio = append(relatorio, <-results)
	}

	return relatorio
}

func workerImportacao(db *gorm.DB, jobs <-chan ImportacaoEscolaDTO, results chan<- ResultadoImportacao) {
	for dados := range jobs {
		err := db.Transaction(func(tx *gorm.DB) error {
			return processarEscola(tx, dados)
		})

		// Retorna o resultado individual no canal
		if err != nil {
			results <- ResultadoImportacao{Inep: dados.Inep, Nome: dados.Nome, Sucesso: false, Erro: err}
		} else {
			results <- ResultadoImportacao{Inep: dados.Inep, Nome: dados.Nome, Sucesso: true, Erro: nil}
		}
	}
}

func processarEscola(tx *gorm.DB, dados ImportacaoEscolaDTO) error {
	var count int64
	if err := tx.Model(&Escola{}).Where("inep = ?", dados.Inep).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return fmt.Errorf("escola com INEP %s já cadastrada", dados.Inep)
	}

	// 2. Instancia a escola e o endereço
	escola, err := NewImportacaoEscola(dados)
	if err != nil {
		return err
	}

	//Salva a escola
	if err := tx.Create(&escola).Error; err != nil {
		return err
	}

	// 3. Salva escolaridades (Pivot)
	escolaridadeMap, err := salvarEscolaridadeEscola(tx, escola.Id, dados.EscolaridadesCodigos)
	if err != nil {
		return err
	}

	turnoMap, err := salvarEscolaTurno(tx, escola.Id, dados.TurnosCodigos)
	if err != nil {
		return err
	}

	//Salva as vagas
	salvarVagasEscolares(tx, escola.Id, escolaridadeMap, turnoMap, dados.Vagas)

	return nil
}

func salvarEscolaridadeEscola(tx *gorm.DB, escolaId string, codigos []string) (map[string]string, error) {
	escolaridadeMap := make(map[string]string)

	if len(codigos) == 0 {
		return escolaridadeMap, nil
	}

	var escolaridades []Escolaridade

	if err := tx.Where("codigo IN ?", codigos).Find(&escolaridades).Error; err != nil {
		return nil, err
	}

	if len(escolaridades) != len(codigos) {
		return nil, fmt.Errorf("Uma ou mais escolaridades não encontrada")
	}

	var pivos []EscolaridadeEscola
	for _, e := range escolaridades {
		escolaridadeMap[e.Codigo] = e.Id
		pivo, err := NewEscolaridadeEscola(escolaId, e.Id)
		if err != nil {
			return nil, err
		}

		pivos = append(pivos, *pivo)
	}

	if err := tx.Create(&pivos).Error; err != nil {
		return nil, err
	}

	return escolaridadeMap, nil
}

func salvarEscolaTurno(tx *gorm.DB, escolaId string, codigos []string) (map[string]string, error) {
	escolaTurnoMap := make(map[string]string)

	if len(codigos) == 0 {
		return escolaTurnoMap, nil
	}

	var turnos []Turno
	if err := tx.Where("codigo IN ?", codigos).Find(&turnos).Error; err != nil {
		return nil, err
	}

	if len(codigos) != len(turnos) {
		return nil, fmt.Errorf("Um ou mais turno não encontrado")
	}

	var pivos []EscolaTurno
	for _, t := range turnos {
		escolaTurnoMap[t.Codigo] = t.Id
		pivo, err := NewEscolaTurno(escolaId, t.Id)
		if err != nil {
			return nil, err
		}

		pivos = append(pivos, *pivo)
	}

	if err := tx.Create(&pivos).Error; err != nil {
		return nil, err
	}

	return escolaTurnoMap, nil
}

func salvarVagasEscolares(tx *gorm.DB, escolaId string, escolaridadeMap, turnoMap map[string]string, vagasDTO []VagasDTO) error {
	if len(vagasDTO) == 0 {
		return nil
	}

	var pivos []VagasEscolares

	for _, vagaDTO := range vagasDTO {
		fmt.Println(vagaDTO.Qtde)
		turnoId := turnoMap[vagaDTO.TurnoCodigo]
		escolaridadeId := escolaridadeMap[vagaDTO.EscolaridadeCodigo]

		if turnoId == "" || escolaridadeId == "" {
			return fmt.Errorf("turno ou escolaridade inválido na vaga")
		}

		pivo, err := NewVagaEscolar(escolaId, escolaridadeId, turnoId, vagaDTO.Qtde, vagaDTO.AnoLetivo)
		if err != nil {
			return err
		}
		pivos = append(pivos, *pivo)
	}

	if len(pivos) > 0 {
		if err := tx.Create(&pivos).Error; err != nil {
			return err
		}
	}

	return nil
}
