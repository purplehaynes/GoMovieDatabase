package handlers

import (
	"GoMovieDB/entities"
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

type Service interface {
	CreateNewMovie(film entities.Movie) error
	ReadAll() (*entities.MVStruct, error)
	GetByMovieId(id string) (*entities.Movie, error)
	DeleteMovieId(id string) error
	UpdateMovieInfo(id string, film entities.Movie) error
}

type MovieHandler struct {
	Svc Service
}

func NewMovieHandler(s Service) MovieHandler {
	return MovieHandler {
		Svc: s,
	}
}

func (mh MovieHandler) CreateMovie(w http.ResponseWriter, r *http.Request) {

	mv := entities.Movie{}

	err := json.NewDecoder(r.Body).Decode(&mv)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = mh.Svc.CreateNewMovie(mv)
	if err != nil {
		switch err.Error() {
		case "movie already exists in database":
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		case "invalid rating":
			http.Error(w, err.Error(), http.StatusNotAcceptable)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

}

func (mh MovieHandler) ReadAll(w http.ResponseWriter, r *http.Request) {

	readDB, err := mh.Svc.ReadAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	readRequest, err := json.MarshalIndent(readDB, "", "	")
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(readRequest)
}

func (mh MovieHandler) GetByMovieId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["Id"]

	getID, err := mh.Svc.GetByMovieId(id)
	if err != nil {
		switch err.Error() {
		case "movie not found":
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
	}

	movieInfo, err := json.MarshalIndent(getID, "", "	")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(movieInfo)
}

func (mh MovieHandler) DeleteMovieId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["Id"]

	err := mh.Svc.DeleteMovieId(id)
	if err != nil {
		switch err.Error() {
		case "movie does not exist":
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (mh MovieHandler) UpdateMovieInfo(w http.ResponseWriter, r *http.Request) {
	mv := entities.Movie{}
	vars := mux.Vars(r)
	id := vars["Id"]

	err := json.NewDecoder(r.Body).Decode(&mv)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = mh.Svc.UpdateMovieInfo(id, mv)
	if err != nil {
		switch err.Error() {
		case "id is mismatched":
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}