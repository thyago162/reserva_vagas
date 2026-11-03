package reservas

type Repository interface {
	Index() ([]*Reserva, error)
	Store(reserva *Reserva) error
	StoreWithTx(tx any, reserva *Reserva) error
	Show(id string) (*Reserva, error)
	Destroy(id string) error
}
