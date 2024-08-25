package server

import (
	"francorutigliano/githubstats/internal/controllers"
	"net/http"
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
	router := http.NewServeMux()

	fs := http.FileServer(http.Dir("./internal/web/static"))
	router.Handle("/static/", http.StripPrefix("/static/", fs))

	router.HandleFunc("POST /user-activity", controllers.UserActivity)

	return http.ListenAndServe(s.port, router)
}
