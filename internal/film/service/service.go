package service

import (
	"database/sql"
	"film-collection/internal/film"
)

type Service struct {
	repo film.Repository
}

func NewService(repo film.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) SaveFilm(film *film.Film) error {
	return s.repo.SaveFilm(film)
}

func (s *Service) UpdateFilm(db *sql.DB, id int, name, discription, releaseDate string, rate int) error {
	return s.repo.UpdateFilm(db, id, name, discription, releaseDate, rate)
}

func (s *Service) DeleteFilm(filmID int) error {
	return s.repo.DeleteFilm(filmID)
}

func (s *Service) AddActorToFilm(filmID, actorID int) error {
	return s.repo.AddActorToFilm(filmID, actorID)
}
func (s *Service) DeleteActorFromFilm(filmID, ActorID int) error {
	return s.repo.DeleteActorFromFilm(filmID, ActorID)
}
