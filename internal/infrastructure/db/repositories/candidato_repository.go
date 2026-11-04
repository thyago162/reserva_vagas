package repositories

import (
	"errors"
	"reserva_vagas/internal/candidato"

	"gorm.io/gorm"
)

type CandidatoRepository struct {
	Db *gorm.DB
}

// NewCandidatoRepository creates a new instance of CandidatoRepository with the provided GORM database connection.
func NewCandidatoRepository(db *gorm.DB) *CandidatoRepository {
	return &CandidatoRepository{
		Db: db,
	}
}

// Index retrieves all candidatos from the database.
func (r *CandidatoRepository) Index() ([]*candidato.Candidato, error) {
	var candidatos []*candidato.Candidato
	tx := r.Db.Find(&candidatos)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return candidatos, nil
}

// Store adds a new candidato to the database.
func (r *CandidatoRepository) Store(candidato *candidato.Candidato) error {
	tx := r.Db.Create(candidato)

	return tx.Error
}

func (r *CandidatoRepository) StoreWithTx(tx any, candidato *candidato.Candidato) error {
	gormTx, ok := tx.(*gorm.DB)
	if !ok {
		return errors.New("transação inválida ou tipo incorreto")
	}

	if err := gormTx.Create(candidato).Error; err != nil {
		return err
	}

	return nil
}

// Show retrieves a candidato by its ID from the database.
func (r *CandidatoRepository) Show(id string) (*candidato.Candidato, error) {
	candidato := &candidato.Candidato{}
	tx := r.Db.Where("id = ?", id).First(candidato)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return candidato, nil
}

// Update modifies an existing candidato in the database.
func (r *CandidatoRepository) Update(candidato *candidato.Candidato) error {
	tx := r.Db.Save(candidato)

	return tx.Error
}

// Delete removes a candidato from the database by its ID.
func (r *CandidatoRepository) Destroy(id string) error {
	var candidato candidato.Candidato
	tx := r.Db.Where("id = ?", id).Delete(&candidato)

	return tx.Error
}
