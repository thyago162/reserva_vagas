package seed

import (
	"fmt"
	"log"
	"math/rand"
	"reserva_vagas/internal/usuario"
	"strings"

	"gorm.io/gorm"
)

func CriarAdminUser(db *gorm.DB) {
	email := "admin@domain.com"
	var adminExiste usuario.Usuario

	db.Where("email = ?", email).Find(&adminExiste)

	if adminExiste.Email != "" {
		return
	}

	role := usuario.Admin
	senhaText := gerarStringAleatoria(10)
	fmt.Printf("Utilize a senha %s para realizar login no sistema, após o primeiro login a senha precisará ser redefinida.", senhaText)
	admin_user, err := usuario.NewUsuario(email, senhaText, role)
	if err != nil {
		log.Fatal("Erro na validação de dados do usuário")
	}

	db.Create(&admin_user)
}

func gerarStringAleatoria(n int) string {
	const letras = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var sb strings.Builder
	for i := 0; i < n; i++ {
		// Escolhe um caractere aleatório do conjunto
		idx := rand.Intn(len(letras))
		sb.WriteByte(letras[idx])
	}
	return sb.String()
}
