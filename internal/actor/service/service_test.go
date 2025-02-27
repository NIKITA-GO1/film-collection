package service

import (
	"database/sql"
	"errors"
	"film-collection/internal/actor"
	"testing"
)

// Мок для actor.Repository
type MockActorRepository struct {
	saveError   error
	updateError error
	deleteError error
}

func (m *MockActorRepository) SaveActor(actor *actor.Actor) error {
	return m.saveError
}

func (m *MockActorRepository) UpdateActor(db *sql.DB, id int, name, gender, birthDate any) error {
	return m.updateError
}

func (m *MockActorRepository) DeleteActor(db *sql.DB, id int) error {
	return m.deleteError
}

func TestSaveActor(t *testing.T) {
	mockRepo := &MockActorRepository{}
	service := NewService(mockRepo)

	actor := &actor.Actor{ID: 1, Name: "John Doe", Gender: "Male", BirthDate: "1990-01-01"}

	// Установка ожидаемого поведения
	mockRepo.saveError = nil

	err := service.SaveActor(actor)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestUpdateActor(t *testing.T) {
	mockRepo := &MockActorRepository{}
	service := NewService(mockRepo)

	db := &sql.DB{} // Здесь можно использовать поддельный объект, если необходимо
	actorID := 1
	name := "John Doe"
	gender := "Male"
	birthDate := "1990-01-01"

	// Установка ожидаемого поведения
	mockRepo.updateError = nil

	err := service.UpdateActor(db, actorID, name, gender, birthDate)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestDeleteActor(t *testing.T) {
	mockRepo := &MockActorRepository{}
	service := NewService(mockRepo)

	db := &sql.DB{} // Здесь можно использовать поддельный объект, если необходимо
	actor := &actor.Actor{ID: 1}

	// Установка ожидаемого поведения
	mockRepo.deleteError = nil

	err := service.DeleteActor(db, actor)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestSaveActor_Error(t *testing.T) {
	mockRepo := &MockActorRepository{}
	service := NewService(mockRepo)

	actor := &actor.Actor{ID: 1, Name: "John Doe", Gender: "Male", BirthDate: "1990-01-01"}

	// Установка ожидаемого поведения
	mockRepo.saveError = errors.New("database error")

	err := service.SaveActor(actor)

	if err == nil {
		t.Errorf("Expected an error, got none")
	} else if err.Error() != "database error" {
		t.Errorf("Expected 'database error', got %v", err)
	}
}
