package main

import (
	"API/Controller"
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
	"log"
	"net/http"
)

func init() {
	gotenv.Load()
}

var db *sql.DB

func main() {

	userLogin, userPassword, dbUrl, passM, token := Controller.ConfigureEnvironmentVar()

	db = Controller.DbCfg(db, dbUrl)

	router := mux.NewRouter()

	router.HandleFunc("/login", Controller.Login(userLogin, userPassword, token)).Methods("POST")
	router.HandleFunc("/register-client", Controller.RegisterClient(db, passM, token))
	router.HandleFunc("/show-clients", Controller.GetClients(db, token)).Methods("GET")
	router.HandleFunc("/show-client", Controller.ShowClient(db, token))
	router.HandleFunc("/delete-client", Controller.DeleteClient(db, passM, token)).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
