package server

import (
	"francorutigliano/githubstats/internal/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	port string
}

func NewServer(p string) *Server {
	return &Server{
		port: p,
	}
}

func (s *Server) Run() error {
	r := mux.NewRouter()

	r.HandleFunc("/", controllers.Home).Methods("GET")
	r.HandleFunc("/search", controllers.Search).Methods("GET")
	r.HandleFunc("/user-activity", controllers.UserActivity).Methods("POST")

	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	log.Printf("server running on localhost%s\n", s.port)
	return http.ListenAndServe(s.port, r)
}
