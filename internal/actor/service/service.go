package service

import (
	"database/sql"
	"film-collection/internal/actor"
)

type Service struct {
	repo actor.Repository
}

func NewService(repo actor.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) SaveActor(actor *actor.Actor) error {
	return s.repo.SaveActor(actor)
}

func (s *Service) UpdateActor(db *sql.DB, id int, name, gender, birthDate any) error {
	return s.repo.UpdateActor(db, id, name, gender, birthDate)
}

func (s *Service) DeleteActor(db *sql.DB, actor *actor.Actor) error {
	return s.repo.DeleteActor(db, actor.ID)
}
