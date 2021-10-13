package handlers

import (
	"github.com/gorilla/mux"
)

type Server struct {
}
	
func ConfigureRouter(handler MovieHandler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/movie", handler.CreateMovie).Methods("POST")
	r.HandleFunc("/movie", handler.ReadAll).Methods("GET")
	r.HandleFunc("/movie/{Id}", handler.GetByMovieId).Methods("GET")
	r.HandleFunc("/movie/{Id}", handler.DeleteMovieId).Methods("DELETE")
	r.HandleFunc("/movie/{Id}", handler.UpdateMovieInfo).Methods("PUT")

	return r
}