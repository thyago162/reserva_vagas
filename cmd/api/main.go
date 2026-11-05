package main

import (
	"log"
	"net/http"
	"reserva_vagas/internal/endpoints"
	"reserva_vagas/internal/importacao"
	"reserva_vagas/internal/infrastructure/db"
	"reserva_vagas/internal/infrastructure/db/seed"

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
	database, err := db.NewDB()
	if err != nil {
		log.Fatal(err)
	}

	//Seed
	seed.Seed(database)

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	endpoints.InitRoutes(r, database)
	importacao.SetDatabase(database)

	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal("Erro ao iniciar o servidor:", err)
	}
}
