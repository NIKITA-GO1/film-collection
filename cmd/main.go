package main

import (
	"database/sql"
	"film-collection/models"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	const connStr = "postgres://user:password@localhost:5432/mydb?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	fmt.Println("Успешное открытие бд")
	query := `drop table if exists films;
			create table if not exists films(
				name text,
				description text,
				release_date text,
				rate integer,
				actors text[]
			);
	`
	_, err = db.Exec(query)
	if err != nil {
		panic(err)
	}
	actors_cast := make([]string, 5)
	actors_cast = append(actors_cast, "jason statham")
	fmt.Println("Create table films")
	meg := models.Film{Name: "meg", Discription: "film pro akuly", ReleaseDate: "hz 2019 god mb", Rate: 10, Actors: actors_cast}
	err = meg.AddFilm(meg)
	err = meg.DeleteActorFromFilm("Jake Chan", meg.Name)
	fmt.Println(err)
}
