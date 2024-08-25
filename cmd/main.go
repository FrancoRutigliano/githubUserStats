package main

import (
	"francorutigliano/githubstats/config"
	"francorutigliano/githubstats/internal/server"
	"log"
)

func main() {
	port, err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}
	s := server.NewServer(port)
	if err := s.Run(); err != nil {
		log.Fatal("error to inicialize the server: ", err)
	}
}
