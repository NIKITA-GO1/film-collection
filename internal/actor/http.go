package actor

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"
)

type ActorHandler struct {
	service Service
}

func NewActorHandler(service Service) *ActorHandler {
	return &ActorHandler{service: service}
}

func (h *ActorHandler) CreateActor(w http.ResponseWriter, r *http.Request) {
	var actor Actor
	err := json.NewDecoder(r.Body).Decode(&actor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if _, err := govalidator.ValidateStruct(actor); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.service.SaveActor(&actor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Println(json.NewEncoder(w).Encode(actor))
}

func (h *ActorHandler) UpdateActor(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Path[len("/actors/"):])
	if err != nil {
		http.Error(w, "Invalid actor ID", http.StatusBadRequest)
		return
	}

	var actor Actor
	err = json.NewDecoder(r.Body).Decode(&actor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if _, err := govalidator.ValidateStruct(actor); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.service.UpdateActor(nil, id, actor.Name, actor.Gender, actor.BirthDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *ActorHandler) DeleteActor(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Path[len("/actors/"):])
	if err != nil {
		http.Error(w, "Invalid actor ID", http.StatusBadRequest)
		return
	}

	err = h.service.DeleteActor(nil, &Actor{ID: id}) // передаем nil, так как db передается в сервис извне
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
