package service

import (
	"GoMovieDB/entities"
	repo "GoMovieDB/repository"
	"errors"

	"github.com/google/uuid"
)

type Service struct {
	Repo repo.Repo
}

func NewService(r repo.Repo) Service {
	return Service{
		Repo: r,
	}
}

func (s Service) CreateNewMovie(film entities.Movie) error {
	film.Id = uuid.New().String()

	if film.Rating >= 0 && film.Rating <= 10 {
		err := s.Repo.CreateNewMovie(film)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s Service) ReadAll() (repo.MVStruct, error) {
	view, err := s.Repo.ReadAll()
	if err != nil {
		return view, errors.New("cannot locate data")
	}
	return view, nil
}


func (s Service) GetByMovieId(id string) (*entities.Movie, error) {
	searchRequest, err := s.Repo.GetMovieId(id)
	if err != nil {
		return nil, err
	}
	return searchRequest, nil
}