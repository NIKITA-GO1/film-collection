package film

import "database/sql"

type Repository interface {
	SaveFilm(film *Film) error
	UpdateFilm(db *sql.DB, id int, name, discription, releaseDate, rate, actors any) error
	DeleteFilm(filmID int) error
	AddActorToFilm(filmID, actorID int) error
	DeleteActorFromFilm(filmID, actorID int) error
}
