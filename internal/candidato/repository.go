package candidato

type Repository interface {
	Index() ([]*Candidato, error)
	Store(candidato *Candidato) error
	StoreWithTx(tx any, candidato *Candidato) error
	Show(id string) (*Candidato, error)
	Update(candidato *Candidato) error
	Destroy(id string) error
	FindByCpf(cpf string) (*Candidato, error)
}
