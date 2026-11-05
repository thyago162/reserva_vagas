package repositories

import (
	"reserva_vagas/internal/usuario"

	"gorm.io/gorm"
)

type UsuarioRepository struct {
	Db *gorm.DB
}

func NewUsuarioRepository(db *gorm.DB) *UsuarioRepository {
	return &UsuarioRepository{Db: db}
}

func (r *UsuarioRepository) Store(usuario *usuario.Usuario) error {
	tx := r.Db.Create(&usuario)

	return tx.Error
}

func (r *UsuarioRepository) FindByEmail(email string) (*usuario.Usuario, error) {
	var usuario *usuario.Usuario
	err := r.Db.Where("email = ? ", email).First(&usuario).Error

	return usuario, err
}
