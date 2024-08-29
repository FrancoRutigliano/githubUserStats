package controllers

import (
	"francorutigliano/githubstats/views"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	views.Home().Render(r.Context(), w)
}
