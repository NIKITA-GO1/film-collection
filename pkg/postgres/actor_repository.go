package postgres

import (
	"database/sql"
	"film-collection/internal/actor"

	"github.com/Masterminds/squirrel"
)

type ActorRepository struct {
	db *sql.DB
}

func NewActorRepository(db *sql.DB) *ActorRepository {
	return &ActorRepository{db: db}
}

func (r *ActorRepository) SaveActor(a *actor.Actor) error {
	query := "INSERT INTO actors(actor_name,gender, birth_date) VALUES ($1, $2, $3) RETURNING id"
	err := r.db.QueryRow(query, a.Name, a.Gender, a.BirthDate).Scan(&a.ID)
	return err
}

func (r *ActorRepository) UpdateActor(db *sql.DB, id int, name, gender, birthDate any) error {
	var finalError error
	queryMap := make(map[string]any)
	queryMap["ID"] = id
	if name != nil {
		queryMap["actor_name"] = name
	}
	if gender != nil {
		queryMap["gender"] = gender
	}
	if birthDate != nil {
		queryMap["birth_date"] = birthDate
	}
	for key, value := range queryMap {
		query, args, err := squirrel.Update("actors").
			Set(key, value).Where(squirrel.Eq{"ID": queryMap["ID"]}).ToSql()

		if err != nil {
			return err
		}
		_, err = db.Exec(query, args...)
		finalError = err

	}
	return finalError
}

func (r *ActorRepository) DeleteActor(db *sql.DB, id int) error {
	query, args, err := squirrel.Delete("actors").Where(squirrel.Eq{"ID": id}).ToSql()
	if err != nil {
		return err
	}
	_, err = db.Exec(query, args...)
	return err
}
