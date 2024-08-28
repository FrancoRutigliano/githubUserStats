package controllers

import (
	"francorutigliano/githubstats/internal/templates"
	"net/http"
)

func Search(w http.ResponseWriter, r *http.Request) {
	c := templates.Search()

	err := templates.Layout(c, "search user").Render(r.Context(), w)
	if err != nil {
		http.Error(w, "error rendering template", http.StatusInternalServerError)
	}
}
