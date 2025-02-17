package postgres

import (
	"database/sql"
	"film-collection/internal/film"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/lib/pq"
)

type FilmRepository struct {
	db *sql.DB
}

func NewFilmRepository(db *sql.DB) *FilmRepository {
	return &FilmRepository{db: db}
}
func (r *FilmRepository) SaveFilm(f *film.Film) error {
	query := "Insert into films(name, discription, release_date, rate, actors_id) Values($1,$2,$3,$4,$5) returning id;"
	err := r.db.QueryRow(query, f.Name, f.Discription, f.ReleaseDate, f.Rate, pq.Array(f.Actors)).Scan(&f.ID)
	return err
}

func (r *FilmRepository) UpdateFilm(db *sql.DB, id int, name, discription, releaseDate string, rate int) error {
	queryBuilder := squirrel.Update("films")

	updateCount := 0
	if name != "" {
		queryBuilder = queryBuilder.Set("name", name)
		updateCount++
	}
	if discription != "" { // Убедитесь, что здесь то же самое имя, что и в БД
		queryBuilder = queryBuilder.Set("discription", discription)
		updateCount++
	}
	if releaseDate != "" {
		queryBuilder = queryBuilder.Set("release_date", releaseDate)
		updateCount++
	}
	if rate != 0 {
		queryBuilder = queryBuilder.Set("rate", rate)
		updateCount++
	}

	// Если ничего не обновляется, возвращаем ошибку
	if updateCount == 0 {
		return fmt.Errorf("no fields to update")
	}

	// Генерация SQL-запроса и аргументов
	query, args, err := queryBuilder.Where("id = ?", id).ToSql()
	if err != nil {
		return err
	}

	fmt.Println("Query:", query)
	fmt.Println("Args:", args)

	// Проверяем, что args имеет правильное количество элементов
	if len(args) != updateCount {
		return fmt.Errorf("argument count mismatch, expected %d but got %d", updateCount, len(args))
	}

	// Выполняем запрос
	_, err = db.Exec(query, args...)
	return err

	/*queryMap := make(map[string]any)
	queryMap["ID"] = id
	if name != nil {
		queryMap["name"] = name
	}
	if discription != nil {
		queryMap["discription"] = discription
	}
	if releaseDate != nil {
		queryMap["releaseDate"] = releaseDate
	}
	if rate != nil {
		queryMap["rate"] = rate
	}
	if actors != nil {
		queryMap["actors"] = actors
	}
	for key, value := range queryMap {
		fmt.Println(key, value)
		query, args, err := squirrel.Update("films").
			Set(key, value).Where(squirrel.Eq{"id": queryMap["ID"]}).ToSql()

		if err != nil {
			return err
		}
		_, err = db.Exec(query, args...)
		finalError = err

	}
	return finalError
	*/
}

func (r *FilmRepository) DeleteFilm(id int) error {
	query := "Delete from films where id = $1"
	_, err := r.db.Exec(query, id)
	return err
}

func (r *FilmRepository) AddActorToFilm(filmID, actorID int) error {
	query := "Update films set actors_id = array_append(actors_id, $2) where id = $1"
	_, err := r.db.Exec(query, filmID, actorID)
	return err
}
func (r *FilmRepository) DeleteActorFromFilm(filmID, actorID int) error {
	query := "Update films set actors_id = array_remove(actors_id, $2) where id = $1"
	_, err := r.db.Exec(query, filmID, actorID)
	return err
}
