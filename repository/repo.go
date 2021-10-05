package repo

import (
	"GoMovieDB/entities"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

type MVStruct struct {
	Movies []entities.Movie
}

type Repo struct {
	Filename string
}

func NewRepository(fn string) Repo {
	return Repo {
		Filename: fn,
	}
}

func (r Repo) CreateNewMovie(film entities.Movie) error {

	mv := MVStruct{}

	output, err := ioutil.ReadFile(r.Filename)
	if err != nil {
		return err
	}

	err = json.Unmarshal(output, &mv)
	if err != nil {
		return err
	}

	for _, v := range mv.Movies { 
		if v.Title == film.Title {
		return errors.New("movie already exists")			
		}
	}

	mv.Movies = append(mv.Movies, film)

	input, err := json.MarshalIndent(mv, "", "	")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(r.Filename, input, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (r Repo) ReadAll() (MVStruct, error) {
	mv := MVStruct{}

	file, err := ioutil.ReadFile(r.Filename)
	if err != nil {
		return mv, errors.New("data field empty")
	}	

	err = json.Unmarshal(file, &mv)

	return mv, err

}

func (r Repo) GetMovieId(id string) (*entities.Movie, error) {

	mv := MVStruct{}

	file, err := ioutil.ReadFile(r.Filename)
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(file, &mv)
	if err != nil {
		return nil, err
	}

	for _, v := range mv.Movies {
		if v.Id == id {
			return &v, nil
		}

	}

	return nil, errors.New("movie not found")

}




	// output, err := ioutil.ReadFile("moviedb.json")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// film.SetId()

	// db := MVStruct{}
	// err = json.Unmarshal(output, &db)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// db.Movies = append(db.Movies, film)


	// input, err := json.MarshalIndent(db, "", "	")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// err = ioutil.WriteFile("moviedb.json", input, 0644)

	// return db, err




