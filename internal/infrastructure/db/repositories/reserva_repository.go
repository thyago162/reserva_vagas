package repositories

import (
	"reserva_vagas/internal/reservas"

	"gorm.io/gorm"
)

type ReservaRepository struct {
	db *gorm.DB
}

func NewReservaRepository(db *gorm.DB) *ReservaRepository {
	return &ReservaRepository{db: db}
}

// Index retrieves all reservas from the database.
func (r *ReservaRepository) Index() ([]*reservas.Reserva, error) {
	var reservas []*reservas.Reserva
	if err := r.db.Find(&reservas).Error; err != nil {
		return nil, err
	}

	return reservas, nil
}

// Store saves a new reserva to the database.
func (r *ReservaRepository) Store(reserva *reservas.Reserva) error {
	if err := r.db.Create(reserva).Error; err != nil {
		return err
	}

	return nil
}

// Show retrieves a reserva by its ID from the database.
func (r *ReservaRepository) Show(id string) (*reservas.Reserva, error) {
	var reserva reservas.Reserva

	if err := r.db.First(&reserva, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &reserva, nil
}

// Destroy removes a reserva from the database.
func (r *ReservaRepository) Destroy(id string) error {
	if err := r.db.Delete(&reservas.Reserva{}, "id = ?", id).Error; err != nil {
		return err
	}

	return nil
}

func (r *ReservaRepository) StoreWithTx(tx *gorm.DB, reserva *reservas.Reserva) error {
	if err := tx.Create(reserva).Error; err != nil {
		return err
	}

	return nil
}
