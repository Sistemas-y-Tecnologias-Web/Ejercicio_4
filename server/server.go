package server

import (
	"log"
	"net/http"

	"videogames-api/internal/videogames"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Start(db *pgxpool.Pool) {
	mux := http.NewServeMux()

	repo := videogames.NewRepository(db)
	service := videogames.NewService(repo)
	handler := videogames.NewHandler(service)
	handler.Register(mux)

	log.Println("API running on :24484")
	log.Fatal(http.ListenAndServe(":24484", mux))
}
