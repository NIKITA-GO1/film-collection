package main

import (
	"database/sql"
	"fmt"

	"film-collection/internal/actor"
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

	actorRepo := postgres.NewActorRepository(db)
	filmRepo := postgres.NewFilmRepository(db)

	actorService := actor.NewService(actorRepo)
	filmService := film.NewService(filmRepo)

	actor1 := &actor.Actor{Name: "Robert Downey Jr.", Gender: "Male", BirthDate: "1965-04-04"}
	actor2 := &actor.Actor{Name: "Chris Evans", Gender: "Male", BirthDate: "1981-06-13"}

	if err := actorService.SaveActor(actor1); err != nil {
		fmt.Println("Error adding actor:", err)
	}
	if err := actorService.SaveActor(actor2); err != nil {
		fmt.Println("Error adding actor:", err)
	}

	// Добавление фильма
	cast := []int{1, 2}
	movie1 := &film.Film{Name: "The Avengers", Discription: "JustFilm", ReleaseDate: "02.10.2012", Rate: 5, Actors: cast}
	if err := filmService.SaveFilm(movie1); err != nil {
		fmt.Println("Error adding film:", err)
	}
	// Назначение актера на фильм
	if err := filmService.AddActorToFilm(movie1.ID, actor1.ID); err != nil {
		fmt.Println("Error adding actor to movie:", err)
	}
	if err := filmService.AddActorToFilm(movie1.ID, actor2.ID); err != nil {
		fmt.Println("Error adding actor to movie:", err)
	}

	// Удаление актера
	if err := filmService.DeleteActorFromFilm(movie1.ID, actor1.ID); err != nil {
		fmt.Println("Error deleting actor:", err)
	}
}
