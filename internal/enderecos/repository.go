package enderecos

type Repository interface {
	Store(endereco *Endereco) error
	Update(endereco *Endereco) error
}
