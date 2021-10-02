package main

import (
	"GoMovieDB/handlers"
	repo "GoMovieDB/repository"
	"GoMovieDB/service"
	"log"
	"net/http"
	"path/filepath"
)

func main() {
	fn := "moviedb.json"

	ext := filepath.Ext(fn)

	if ext != ".json" {
		log.Fatalln("File extension invalid")
	}

	r := repo.NewRepository(fn)

	svc := service.NewService(r)

	hdlr := handlers.NewMovieHandler(svc)

	router := handlers.ConfigureRouter(hdlr)

	svr := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: router,
	}

	log.Fatalln(svr.ListenAndServe())
}
