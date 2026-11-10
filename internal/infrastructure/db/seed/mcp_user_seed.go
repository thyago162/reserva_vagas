package seed

import (
	"log"
	"os"
	"reserva_vagas/internal/usuario"

	"gorm.io/gorm"
)

func CriarMCPUser(db *gorm.DB) {
	email := "mcp@domain.com"
	var mcpExiste usuario.Usuario

	db.Where("email = ?", email).Find(&mcpExiste)

	if mcpExiste.Email != "" {
		return
	}

	role := usuario.Mcp
	senhaText := os.Getenv("MCP_PWD")
	mcp_user, err := usuario.NewUsuario(email, senhaText, role)
	if err != nil {
		log.Fatal("Erro na validação de dados do usuário")
	}

	db.Create(&mcp_user)
}
