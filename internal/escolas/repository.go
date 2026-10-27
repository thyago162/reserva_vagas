package escolas

import "reserva_vagas/internal/importacao"

type Repository interface {
	Index() ([]*Escola, error)
	Store(escola *Escola) error
	Show(id string) (*Escola, error)
	Update(escola *Escola) error
	Destroy(id string) error
	ImportarEmLotes(escolas []importacao.Escola) error
}
