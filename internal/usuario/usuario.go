package usuario

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type TipoRole string

const (
	Admin  TipoRole = "ADMIN"
	Escola TipoRole = "ESCOLA"
	User   TipoRole = "USER"
)

type Usuario struct {
	Id        string    `gorm:"primaryKey;size:36"`
	Email     string    `gorm:"unique;not null"`
	Senha     string    `gorm:"not null"`
	Role      TipoRole  `gorm:"type:enum('ADMIN', 'ESCOLA', 'USER');default:'USER'"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func NewUsuario(email, senha string, role TipoRole) (*Usuario, error) {
	bytesHash, err := bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &Usuario{
		Id:    uuid.NewString(),
		Email: email,
		Senha: string(bytesHash),
		Role:  role,
	}, nil
}
