package Controller

import (
	"API/Model"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func Login(user1 string, url string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		login := r.URL.Query()["login"][0]
		password := r.URL.Query()["password"][0]
		user := login + " " + password
		fmt.Println(user)
		if user == user1 {
			http.Redirect(w, r, url+"logged.html", http.StatusSeeOther)
		} else {
			http.Redirect(w, r, "http://localhost:3000/", http.StatusSeeOther)
		}
	}
}

func RegisterClient(db *sql.DB, url string, passM string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		nameR := r.URL.Query()["name"][0]
		snameR := r.URL.Query()["surname"][0]
		snR := r.URL.Query()["sn"][0]
		passMR := r.URL.Query()["passM"][0]

		if passMR == passM {
			var client Model.Client

			client.Name = nameR
			client.Surname = snameR
			client.Sn = snR

			json.NewDecoder(r.Body).Decode(&client)

			_ = db.QueryRow("insert into clients (name, sname, sn) values($1, $2, $3);", client.Name, client.Surname, client.Sn)
			http.Redirect(w, r, url+"logged.html", http.StatusSeeOther)
		} else {
			http.Redirect(w, r, url+"logged.html?error=wrong-manager-pass", http.StatusSeeOther)
		}
	}
}

func GetClients(db *sql.DB, url string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var client Model.Client
		rows, _ := db.Query("select * from clients")
		defer rows.Close()
		i := 1
		for rows.Next() {
			rows.Scan(&client.Name, &client.Surname, &client.Sn)
			is := strconv.Itoa(i)
			w.Header().Add("Content-Type", "html")
			fmt.Fprintf(w, " "+"<b>"+is+") <br> </b>"+`<b style="margin-left: 30px;"> Client name: </b>`+client.Name+`<br> <b style="margin-left: 30px;"> Client surname: </b>`+client.Surname+`<br> <b style="margin-left: 30px;"> Client S/N: </b>`+client.Sn+"<br>")
			i++
		}
	}
}

func DeleteClient(db *sql.DB, url string, passMD string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		snD := r.URL.Query()["sn"][0]
		passM := r.URL.Query()["passM"][0]
		if passM == passMD {
			result, _ := db.Exec("delete from clients where sn=$1", snD)
			rowsUpdated, _ := result.RowsAffected()
			fmt.Println(rowsUpdated)
			http.Redirect(w, r, url+"logged.html", http.StatusSeeOther)
		} else {
			http.Redirect(w, r, url+"logged.html?error=wrong-manager-passD", http.StatusSeeOther)
		}
	}
}
