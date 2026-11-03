package enderecos

type Repository interface {
	Store(endereco *Endereco) error
	StoreWithTx(tx any, endereco *Endereco) error
	Update(endereco *Endereco) error
}
