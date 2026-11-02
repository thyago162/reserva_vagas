package reservas

type Repository interface {
	Index() ([]*Reserva, error)
	Store(reserva *Reserva) error
	Show(id string) (*Reserva, error)
	Update(reserva *Reserva) error
	Destroy(id string) error
}
