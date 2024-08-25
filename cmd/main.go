package main

import (
	"fmt"
	"francorutigliano/githubstats/config"
	"log"
)

func main() {
	port, err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(port)
}
