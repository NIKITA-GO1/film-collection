package actor

import "database/sql"

type Repository interface {
	SaveActor(actor *Actor) error
	UpdateActor(db *sql.DB, id int, name, gender, birthDate any) error
	DeleteActor(db *sql.DB, id int) error
}
