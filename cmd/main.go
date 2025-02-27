package main

import (
	"database/sql"
	"film-collection/internal/film"
	"film-collection/internal/film/service"
	"film-collection/pkg/postgres"
	"fmt"
	"net/http"

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

	filmService := service.NewService(filmRepo)
	filmHandler := film.NewFilmHandler(*filmService)
	http.HandleFunc("/films", filmHandler.CreateFilm)
	http.HandleFunc("/films/", filmHandler.UpdateFilm)

	err = http.ListenAndServe(`:8080`, nil)
	if err != nil {
		fmt.Println(err)
	}

}
