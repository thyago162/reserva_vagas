package escolas

import "github.com/google/uuid"

type Escola struct {
	Id    string `gorm:"primaryKey"`
	Nome  string
	Inep  string
	Email string
	Sigla string
}

func NewEscola(nome string, inep string, email string, sigla string) (*Escola, error) {
	return &Escola{
		Id:    uuid.NewString(),
		Nome:  nome,
		Inep:  inep,
		Email: email,
		Sigla: sigla,
	}, nil
}
