package escolaridades

type Repository interface {
	Show(escolaridadeId string) (*Escolaridade, error)
}
