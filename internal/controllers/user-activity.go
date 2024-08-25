package controllers

import (
	"encoding/json"
	"fmt"
	"francorutigliano/githubstats/internal/models"
	"net/http"
)

func UserActivity(w http.ResponseWriter, r *http.Request) {

	user := r.URL.Query().Get("user")

	url := fmt.Sprintf("https://api.github.com/users/%s/events?per_page=10", user)

	response, err := http.Get(url)
	if err != nil {
		http.Error(w, "error on request", http.StatusInternalServerError)
		return
	}

	if response.StatusCode == http.StatusNotFound {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	if response.StatusCode != http.StatusOK {
		http.Error(w, "error fetching data", http.StatusInternalServerError)
		return
	}

	var payload []models.GithubEventResponse
	if err = json.NewDecoder(response.Body).Decode(&payload); err != nil {
		http.Error(w, "error serializing the json", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(&payload); err != nil {
		http.Error(w, "error encoding", http.StatusInternalServerError)
		return
	}

}
