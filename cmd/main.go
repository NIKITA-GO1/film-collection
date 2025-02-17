package main

import (
	"database/sql"
	"fmt"

	//"film-collection/internal/actor"
	"film-collection/internal/film"
	"film-collection/pkg/postgres"

	_ "github.com/lib/pq"
)

func main() {
	const connStr = "postgres://user:password@localhost:5432/mydb?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	filmRepo := postgres.NewFilmRepository(db)

	filmService := film.NewService(filmRepo)

	if err := filmService.DeleteActorFromFilm(8, 1); err != nil {
		fmt.Println("Error update film:", err)
	}
}
