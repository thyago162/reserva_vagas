package auth

import (
	"os"
	"time"

	"reserva_vagas/internal/usuario"

	"github.com/golang-jwt/jwt/v5"
)

// Chave secreta usada para assinar o token (em produção, use variável de ambiente)
var jwtSecret = []byte(os.Getenv("SECRECT_KEY"))

// Claims define o que vai dentro do token (Payload)
type Claims struct {
	UserID string           `json:"user_id"`
	Email  string           `json:"email"`
	Role   usuario.TipoRole `json:"role"`
	jwt.RegisteredClaims
}

// GerarToken cria um JWT válido por 24 horas para o usuário autenticado
func GerarToken(u *usuario.Usuario) (string, error) {
	claims := Claims{
		UserID: u.Id,
		Email:  u.Email,
		Role:   u.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // Validade de 1 dia
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	// Cria o token com o método de assinatura HS256 e os dados (claims)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Assina o token usando a chave secreta do servidor
	return token.SignedString(jwtSecret)
}
