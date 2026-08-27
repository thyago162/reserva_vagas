package main

import (
	"log"
	"net/http"
	"reserva_vagas/internal/endpoints"
	"reserva_vagas/internal/escolas"
	"reserva_vagas/internal/infrastructure/db"
	"reserva_vagas/internal/infrastructure/db/repositories"
	"reserva_vagas/internal/infrastructure/db/seed"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()
	database, err := db.NewDB()
	if err != nil {
		log.Fatal(err)
	}

	//Seed
	seed.Seed(database)

	repo := repositories.NewEscolaRepositoryDb(database)
	service := escolas.NewService(repo)
	handler := endpoints.NewEscolaHandler(service)

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/escolas", func(r chi.Router) {
		r.Post("/", endpoints.HandlerErro(handler.Criar))
		r.Get("/", endpoints.HandlerErro(handler.Index))
		r.Get("/{id}", endpoints.HandlerErro(handler.Show))
		r.Put("/{id}", endpoints.HandlerErro(handler.Update))
		r.Delete("/{id}", endpoints.HandlerErro(handler.Destroy))
	})

	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal("Erro ao iniciar o servidor:", err)
	}
}
