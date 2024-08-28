package controllers

import (
	"francorutigliano/githubstats/pkg/utils"
	"net/http"
)

func Search(w http.ResponseWriter, r *http.Request) {
	utils.Render(w, "search.html", "hola mundo")
}
