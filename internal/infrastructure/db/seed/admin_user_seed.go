package seed

import (
	"log"
	"os"
	"reserva_vagas/internal/usuario"

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
	senhaText := os.Getenv("ADMIN_PWD")
	admin_user, err := usuario.NewUsuario(email, senhaText, role)
	if err != nil {
		log.Fatal("Erro na validação de dados do usuário")
	}

	db.Create(&admin_user)
}
