package Controller

import (
	"fmt"
	"net/http"
)

func Login(user1 string, token string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		login := r.URL.Query()["login"][0]
		password := r.URL.Query()["password"][0]
		user := login + " " + password
		fmt.Println(user)
		if user == user1 {
			http.Redirect(w, r, "http://localhost:3000/logged.html"+token, http.StatusSeeOther)
		} else {
			http.Redirect(w, r, "http://localhost:3000/", http.StatusSeeOther)
		}
	}
}
