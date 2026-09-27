package escolas

type EscolaridadeEscolaRepository interface {
	Index(escolaId string) ([]*EscolaridadeEscola, error)
	Store([]*EscolaridadeEscola) error
	Exists(escolaridadeID, escolaID string) (bool, error)
	Destroy(ids []string) error
}
