package actor

import "database/sql"

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) SaveActor(actor *Actor) error {
	return s.repo.SaveActor(actor)
}

func (s *Service) UpdateActor(db *sql.DB, id int, name, gender, birthDate any) error {
	return s.repo.UpdateActor(db, id, name, gender, birthDate)
}

func (s *Service) DeleteActor(db *sql.DB, actor *Actor) error {
	return s.repo.DeleteActor(db, actor.ID)
}
