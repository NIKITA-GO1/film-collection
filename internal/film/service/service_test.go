package service

import (
	"database/sql"
	"film-collection/internal/film"
	"testing"
)

type MockFilmRepository struct {
	saveError   error
	updateError error
	deleteError error
}

func (m *MockFilmRepository) saveFilm(film *film.Film) error {
	return m.saveError
}

func (m *MockFilmRepository) updateFilm(film *film.Film) error {
	return m.updateError
}

func (m *MockFilmRepository) deleteFilm(film *film.Film) error {
	return m.deleteError
}

func TestSaveFilm(t *testing.T) {
	mockRepo := &MockFilmRepository{}
	service := NewService(mockRepo)

	film := &film.Film{Name: "Hunt", Discription: "film about hunter", ReleaseDate: "22.02.1990", Rate: 8, Actors: []int{1, 2}}

	mockRepo.saveError = nil

	err := service.SaveFilm(film)

	if err != nil {
		t.Errorf("new error:", err)
	}
}

func TestUpdateFilm(t *testing.T) {
	mockRepo := &MockFilmRepository{}
	service := NewService(mockRepo)

	film := &film.Film{Name: "proger", Discription: "film about it", ReleaseDate: "22.02.1995", Rate: 5, Actors: []int{4, 5}}
	const connStr = "postgres://user:password@localhost:5432/mydb?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		t.Errorf("error open db", err)
	}
	err = service.UpdateFilm(db, 5, film.Name, film.Discription, film.ReleaseDate, 5)
	if err != nil {
		t.Errorf("error update films:", err)
	}

}

func TestDeleteFilm(t *testing.T) {
	mockRepo := &MockFilmRepository{}
	service := NewService(mockRepo)

	err := service.DeleteFilm(5)
	if err != nil {
		t.Errorf("error while delete film", err)
	}
}
