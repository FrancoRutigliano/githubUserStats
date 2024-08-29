package controllers

import (
	"francorutigliano/githubstats/views"
	"net/http"
)

func Search(w http.ResponseWriter, r *http.Request) {
	views.Search().Render(r.Context(), w)
}
