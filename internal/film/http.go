package film

import (
	"encoding/json"
	"net/http"
	"strconv"
	"unicode/utf8"
)

type FilmHandler struct {
	service Service
}

func NewFilmHandler(service Service) *FilmHandler {
	return &FilmHandler{service: service}
}

func (h *FilmHandler) AddActorToFilm(w http.ResponseWriter, r *http.Request) {
	filmID, err := strconv.Atoi(r.URL.Path[len("/films/"):])
	if err != nil {
		http.Error(w, "Invalid film ID", http.StatusBadRequest)
		return
	}
	var data struct{ ActorID int }
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.service.AddActorToFilm(filmID, data.ActorID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *FilmHandler) DeleteActorFromFilm(w http.ResponseWriter, r *http.Request) {
	filmID, err := strconv.Atoi(r.URL.Path[len("/films/"):])
	if err != nil {
		http.Error(w, "Invalid film ID", http.StatusBadRequest)
		return
	}
	var data struct{ ActorID int }
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.service.DeleteActorFromFilm(filmID, data.ActorID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *FilmHandler) CreateFilm(w http.ResponseWriter, r *http.Request) {
	var film Film
	err := json.NewDecoder(r.Body).Decode(&film)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if utf8.RuneCountInString(film.Name) < 1 || 150 < utf8.RuneCountInString(film.Name) {
		http.Error(w, "Film name must be between 1 and 150 characters", http.StatusBadRequest)
		return
	}
	if utf8.RuneCountInString(film.Discription) > 1000 {
		http.Error(w, "Film description cant be more than 1000 characters", http.StatusBadRequest)
		return
	}
	if film.Rate < 1 || film.Rate > 10 {
		http.Error(w, "Rate must be between 1 and 10", http.StatusBadRequest)
		return
	}

	err = h.service.SaveFilm(&film)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(film)
}

func (h *FilmHandler) UpdateFilm(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Path[len("/films/"):])
	if err != nil {
		http.Error(w, "Invalid film ID", http.StatusBadRequest)
		return
	}

	var film Film
	err = json.NewDecoder(r.Body).Decode(&film)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if utf8.RuneCountInString(film.Name) < 1 || 150 < utf8.RuneCountInString(film.Name) {
		http.Error(w, "Film name must be between 1 and 150 characters", http.StatusBadRequest)
		return
	}
	if utf8.RuneCountInString(film.Discription) > 1000 {
		http.Error(w, "Film description cant be more than 1000 characters", http.StatusBadRequest)
		return
	}
	if film.Rate < 1 || film.Rate > 10 {
		http.Error(w, "Rate must be between 1 and 10", http.StatusBadRequest)
		return
	}

	err = h.service.UpdateFilm(nil, id, film.Name, film.Discription, film.ReleaseDate, film.Rate) // передаем nil, так как db передается в сервис извне
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *FilmHandler) DeleteFilm(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Path[len("/films/"):])
	if err != nil {
		http.Error(w, "Invalid film ID", http.StatusBadRequest)
		return
	}

	err = h.service.DeleteFilm(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
