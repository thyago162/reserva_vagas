package vagasescolares

type Repository interface {
	Index(escolaId string, anoLetivo int) ([]*VagasEscolares, error)
	Store(vagasEscolares []*VagasEscolares) error
	Show(id string) (*VagasEscolares, error)
	Update(vagaEscolar *VagasEscolares) error
	DebitarVaga(tx any, id string) error
}
