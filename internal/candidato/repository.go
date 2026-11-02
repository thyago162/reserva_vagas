package candidato

type Repository interface {
	Index() ([]*Candidato, error)
	Store(candidato *Candidato) error
	Show(id string) (*Candidato, error)
	Update(candidato *Candidato) error
	Delete(id string) error
}
