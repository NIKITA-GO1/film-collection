package film

import "database/sql"

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) SaveFilm(film *Film) error {
	return s.repo.SaveFilm(film)
}

func (s *Service) UpdateFilm(db *sql.DB, id int, name, discription, releaseDate, rate, actors any) error {
	return s.repo.UpdateFilm(db, id, name, discription, releaseDate, rate, actors)
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
