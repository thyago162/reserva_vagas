package vagasescolares

type Repository interface {
	Index(escolaId string) (*[]VagasEscolares, error)
	Store(vagasEscolares []*VagasEscolares) error
	Show(escolaId string) (*VagasEscolares, error)
	Update(vagasEscolares *VagasEscolares) error
	Delete(id string) error
}

//07192784494
