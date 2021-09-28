package repo

import (
	"GoMovieDB/entities"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type MVStruct struct {
	Movies []entities.Movie
}

// func (moov *MVStruct) PostToDb (movie entities.Movie) {
// 	moov.Movies = append(moov.Movies, movie)
// }

func CreateNewMovie (film entities.Movie) (MVStruct, error) {

	output, err := ioutil.ReadFile("moviedb.json")
	if err != nil {
		fmt.Println(err)
	}

	film.SetId()

	db := MVStruct{}
	err = json.Unmarshal(output, &db)
	if err != nil {
		fmt.Println(err)
	}

	db.Movies = append(db.Movies, film)


	input, err := json.MarshalIndent(db, "", "	")
	if err != nil {
		fmt.Println(err)
	}

	err = ioutil.WriteFile("moviedb.json", input, 0644)

	return db, err

}


// type FileConverter struct {
// 	filename string
// }


// func NewFileConverter(fn string) FileConverter {
// 	return FileConverter {
// 		filename: fn,
// 	}
// }

// func (f FileConverter) ConvertToGo() (entities.Movie, error) {
// 	db := entities.Movie{}
// 	jsonBytes, err := ioutil.ReadFile(f.filename)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	err = json.Unmarshal(jsonBytes, &db)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	return db, err

// }

// func (f FileConverter) ConvertToDB(Movie entities.Movie) error {
// 	jsonBytes, err := json.MarshalIndent(Movie, " ", "",)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	err = ioutil.WriteFile(f.filename, jsonBytes, 0644)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	return err
// }