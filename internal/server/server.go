package server

type Server struct {
	port string
}

func NewServer(p string) *Server {
	return &Server{
		port: p,
	}
}
