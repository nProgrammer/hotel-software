package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/hello", hello)

	log.Fatal(http.ListenAndServe(":8000", router))
}

func hello(w http.ResponseWriter, r *http.Request) {
	login := r.URL.Query()["login"][0]
	password := r.URL.Query()["password"][0]
	user := login + " " + password
	w.Header().Add("Content-Type", "text/plain")
	fmt.Fprintf(w, user)
}
