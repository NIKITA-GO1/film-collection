package models

import (
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

type Film struct {
	Name        string
	Discription string
	ReleaseDate string
	Rate        int
	Actors      []string
}

// метод добавлени фильма
func (f *Film) AddFilm(film Film) error {
	if len(film.Name) < 1 || 150 < len(film.Name) || 1000 < len(film.Discription) || film.Rate < 0 || 10 < film.Rate {
		return errors.New("Некорректный размер данных фильма")
	}
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	query := `insert into films(name, description, release_date, rate, actors)
			values($1,$2,$3,$4,$5);`

	_, err = db.Exec(query, film.Name, film.Discription, film.ReleaseDate, film.Rate, pq.Array(film.Actors))
	if err != nil {
		return err
	}
	return nil
}

// метод редактирования названия фильма
func (f *Film) EditFilmName(newName, oldName string) error {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	query := `update films set name = $1 where name = $2;`

	_, err = db.Exec(query, newName, oldName)
	if err != nil {
		return err
	}
	return nil
}

// метод редактирования описания фильма
func (f *Film) EditFilmDescription(newDescription, oldDescription string) error {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	query := `update  films set description = $1 where description = $2;`

	_, err = db.Exec(query, newDescription, oldDescription)
	if err != nil {
		return err
	}
	return nil
}

// метод редактирования даты выхода фильма
func (f *Film) EditFilmDate(newDate, oldDate string) error {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	query := `update films  set release_date = $1 where release_date = $2;`

	_, err = db.Exec(query, newDate, oldDate)
	if err != nil {
		return err
	}
	return nil
}

// метод редактирования оценки фильма
func (f *Film) EditFilmRate(newRate, oldRate string) error {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	query := `update films  set rate = $1 where rate = $2;`

	_, err = db.Exec(query, newRate, oldRate)
	if err != nil {
		return err
	}
	return nil
}

// удаление фильма
func (f *Film) DeleteFilm(name string) error {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	query := `delete from films where name = $1;`

	_, err = db.Exec(query, name)
	if err != nil {
		return err
	}
	return nil
}

// добавление актера к фильму
func (f *Film) AddActorToFilm(actorName, filmName string) error {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	query := `
        UPDATE films
        SET actors = array_append(actors, $1)
        WHERE name = $2;`

	_, err = db.Exec(query, actorName, filmName)
	if err != nil {
		return err
	}
	return nil
}

// удаление актера из фильма
func (f *Film) DeleteActorFromFilm(actorName, filmName string) error {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	query := `
        UPDATE films
        SET actors = array_remove(actors, $1)
        WHERE name = $2;`

	_, err = db.Exec(query, actorName, filmName)
	if err != nil {
		return err
	}
	return nil
}
