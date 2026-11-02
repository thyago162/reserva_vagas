package vagasescolares

type Repository interface {
	Index(escolaId string, anoLetivo int) ([]*VagasEscolares, error)
	Store(vagasEscolares []*VagasEscolares) error
	Show(id string) (*VagasEscolares, error)
	Update(vagaEscolar *VagasEscolares) error
}

//07192784494
