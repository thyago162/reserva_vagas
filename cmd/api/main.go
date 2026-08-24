package main

import (
	"log"
	"net/http"
	"reserva_vagas/internal/endpoints"
	"reserva_vagas/internal/escolas"
	"reserva_vagas/internal/infrastructure/db"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()
	database, err := db.NewDB()
	if err != nil {
		log.Fatal(err)
	}

	repo := db.NewEscolaRepositoryDb(database)
	service := escolas.NewService(repo)
	handler := endpoints.NewEscolaHandler(service)

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/escolas", func(r chi.Router) {
		r.Post("/", handler.Criar)
		r.Get("/", handler.Index)
	})

	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal("Erro ao iniciar o servidor:", err)
	}
}
