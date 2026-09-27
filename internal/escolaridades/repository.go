package escolaridades

type Repository interface {
	Show(escolaridadeId string) (*Escolaridade, error)
	FindByIds(escolaridadeIds []string) (int64, error)
}
