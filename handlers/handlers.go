package handlers

import (
	"GoMovieDB/entities"
	"GoMovieDB/service"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type MovieHandler struct {
	Svc service.Service
}

func NewMovieHandler(s service.Service) MovieHandler {
	return MovieHandler {
		Svc: s,
	}
}


func (mh MovieHandler) PostMovieHandler(w http.ResponseWriter, r *http.Request) {
	mv := entities.Movie{}

	err := json.NewDecoder(r.Body).Decode(&mv)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = mh.Svc.CreateNewMovie(mv)
	if err != nil {
		switch err.Error() {
		case "movie already exists":
			http.Error(w, err.Error(), http.StatusBadRequest)
		case "invalid rating":
			http.Error(w, err.Error(), http.StatusNotAcceptable)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)


}

func (mh MovieHandler) ReadAll(w http.ResponseWriter, r *http.Request) {
	readDB, err := mh.Svc.ReadAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	readRequest, err := json.MarshalIndent(readDB, "", "	")
	if err != nil {
		fmt.Println(err)
	}
	

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(readRequest)
}



func (mh MovieHandler) GetMovieId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["Id"]

	getID, err := mh.Svc.GetMovieId(id)
	if err != nil {
		switch err.Error() {
		case "movie not found":
			http.Error(w, err.Error(), http.StatusNotFound)
		}
	}

	foundID, err := json.MarshalIndent(getID, "", "	")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(foundID)
}







// 	mv := entities.Movie{}
	
// 	getRequest, _ := json.Marshal(mv.Id)

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(getRequest)
// }