package entities

import (
	"github.com/google/uuid"
)

type Movie struct {
	Id 				string
	Title 			string
	Genre 			[]string
	Description 	string
	Director 		string
	Actors 			[]string
	Rating 			float32
}

type MVStruct struct {
	Movies []Movie
}

func (mv *Movie) SetId() {
	mv.Id = uuid.New().String()
}