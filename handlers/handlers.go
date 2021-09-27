package handlers

import (
	"GoMovieDB/entities"
	"GoMovieDB/repo"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func PostNewMovie(w http.ResponseWriter, r *http.Request) {
	mv := entities.Movie{}

	err := json.NewDecoder(r.Body).Decode(&mv)
	if err != nil {
		fmt.Println(err)
	}

	mv.SetId()

	db := repo.MVStruct{}
	db.PostToDb(mv)
	
	jsonBytes, err := json.MarshalIndent(db, "", "	")
	if err != nil {
		fmt.Println(err)
	}

	err = ioutil.WriteFile("moviedb.json", jsonBytes, 0644)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)

}