package usuario

type Repository interface {
	Store(usuario *Usuario) error
	FindByEmail(email string) (*Usuario, error)
}
