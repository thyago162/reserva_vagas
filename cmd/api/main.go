package main

import (
	"log"
	"net/http"
	"reserva_vagas/internal/endpoints"
	"reserva_vagas/internal/importacao"
	"reserva_vagas/internal/infrastructure/db"
	"reserva_vagas/internal/infrastructure/db/seed"
	"reserva_vagas/internal/infrastructure/messasing"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}

	r := chi.NewRouter()

	//Inicializa o banco de dados
	database, err := db.NewDB()
	if err != nil {
		log.Fatal(err)
	}

	//Inicializa o RabbitMQ
	rabbitConn, err := messasing.NewRabbitMQ()
	if err != nil {
		log.Fatal(err)
	}

	defer rabbitConn.Conn.Close()
	defer rabbitConn.Channel.Close()

	publisher := messasing.NewRabbitPublisher(rabbitConn.Channel)

	//Seed
	seed.Seed(database)

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	endpoints.InitRoutes(r, database, publisher)
	importacao.SetDatabase(database)

	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal("Erro ao iniciar o servidor:", err)
	}
}
