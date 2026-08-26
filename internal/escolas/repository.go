package escolas

type Repository interface {
	Salvar(escola *Escola) error
	GetAll() ([]*Escola, error)
	GetById(id string) (*Escola, error)
}
