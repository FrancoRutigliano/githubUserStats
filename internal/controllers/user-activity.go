package controllers

import (
	"net/http"
)

func UserActivity(w http.ResponseWriter, r *http.Request) {
	user := r.URL.Query().Get("user")
	if user == "" {
		http.Error(w, "username is requiered", http.StatusBadRequest)
	}

}
