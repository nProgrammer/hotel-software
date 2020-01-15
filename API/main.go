package main

import (
	"API/Controller"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
	"log"
	"net/http"
	"os"
)

func init() {
	gotenv.Load()
}

func main() {
	userLogin := os.Getenv("USER_LOGIN")
	fmt.Println(userLogin)
	userPassword := os.Getenv("USER_PASSWORD")
	fmt.Println(userPassword)

	user := userLogin + " " + userPassword

	router := mux.NewRouter()

	router.HandleFunc("/login", Controller.Login(user))

	log.Fatal(http.ListenAndServe(":8000", router))
}


