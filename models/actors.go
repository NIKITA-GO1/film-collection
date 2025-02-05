package models

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Actor struct {
	Name      string
	Gender    string
	BirthDate string
}

const connStr = "postgres://user:password@localhost:5432/mydb?sslmode=disable"

// Добавление актера в систему
func (a *Actor) CreateActor(name, gender, birthDate string) error {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	query := `insert into actors(gender, actor_name, birth_date)
			values($1,$2,$3)
	`
	_, err = db.Exec(query, name, gender, birthDate)
	if err != nil {
		return err
	}
	return nil
}

// метод редактирования имени актера
func (a *Actor) EditActorName(lastName, newName string) error {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	query := `update actors set actor_name = $1 where actor_name = $2;`
	_, err = db.Exec(query, newName, lastName)
	if err != nil {
		return err
	}
	return nil
}

// метод редактирования пола актера
func (a *Actor) EditActorGender(newGender, name string) error {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	query := `update actors set gender = $1 where actor_name = $2;`
	_, err = db.Exec(query, newGender, name)
	if err != nil {
		return err
	}
	return nil
}

// метод редактирования даты рождения актера
func (a *Actor) EditActorBirthDate(newBirthDate, name string) error {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	query := `update actors set birth_date = $1 where actor_name = $2;`
	_, err = db.Exec(query, newBirthDate, name)
	if err != nil {
		return err
	}
	return nil
}

func (a *Actor) DeleteActor(name string) error {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	query := `delete from actors where actor_name = $1`
	_, err = db.Exec(query, name)
	if err != nil {
		return err
	}
	return nil

}
