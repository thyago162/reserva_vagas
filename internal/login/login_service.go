package login

import (
	"fmt"
	"reserva_vagas/internal/usuario"

	"golang.org/x/crypto/bcrypt"
)

type LoginService struct {
	repo usuario.Repository
}

func NewLoginService(r usuario.Repository) *LoginService {
	return &LoginService{repo: r}
}

func (s *LoginService) Login(email, senha string) (*usuario.Usuario, error) {
	fmt.Println(email)
	usuario, err := s.repo.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(usuario.Senha), []byte(senha))
	if err != nil {
		return nil, err
	}

	return usuario, nil
}
