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

	router.HandleFunc("POST /user-activity", controllers.UserActivity)

	return http.ListenAndServe(s.port, router)
}
