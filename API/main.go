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

	userLogin, userPassword, url, dbUrl, passM, token, aUrl := Controller.ConfigureEnvironmentVar()

	db = Controller.DbCfg(db, dbUrl)

	user := userLogin + " " + userPassword

	router := mux.NewRouter()

	router.HandleFunc("/login", Controller.Login(user, url, token))
	router.HandleFunc("/register-client", Controller.RegisterClient(db, url, passM))
	router.HandleFunc("/show-clients", Controller.GetClients(db, aUrl))
	router.HandleFunc("/show-client", Controller.ShowClient(db))
	router.HandleFunc("/delete-client", Controller.DeleteClient(db, url, passM))

	log.Fatal(http.ListenAndServe(":8000", router))

}
