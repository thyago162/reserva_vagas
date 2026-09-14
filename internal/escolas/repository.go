package escolas

type Repository interface {
	Index() ([]*Escola, error)
	Store(escola *Escola) error
	Show(id string) (*Escola, error)
	Update(escola *Escola) error
	Destroy(id string) error
}
