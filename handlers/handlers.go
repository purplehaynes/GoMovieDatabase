package handlers

import (
	"GoMovieDB/entities"
	"GoMovieDB/repo"
	"encoding/json"
	"fmt"
	"net/http"
)

func PostNewMovie(w http.ResponseWriter, r *http.Request) {
	mv := entities.Movie{}

	err := json.NewDecoder(r.Body).Decode(&mv)
	if err != nil {
		fmt.Println(err)
	}

	addMovie, err := repo.CreateNewMovie(mv)
	if err != nil {
		fmt.Println(err)
	}

	
	jsonBytes, err := json.MarshalIndent(addMovie, "", "	")
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)

}