package Controller

import (
	"API/Model"
	"API/View"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func Login(user1 string, url string, token string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		login := r.URL.Query()["login"][0]
		password := r.URL.Query()["password"][0]
		tokenU := r.URL.Query()["token"][0]
		user := login + " " + password
		fmt.Println(user)
		if user == user1 && tokenU == token {
			http.Redirect(w, r, url+"logged.html", http.StatusSeeOther)
		} else {
			http.Redirect(w, r, url, http.StatusSeeOther)
		}
	}
}

func RegisterClient(db *sql.DB, url string, passM string, token string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		nameR := r.URL.Query()["name"][0]
		snameR := r.URL.Query()["surname"][0]
		snR := r.URL.Query()["sn"][0]
		roomNumber := r.URL.Query()["roomN"][0]
		carId := r.URL.Query()["carId"][0]
		passMR := r.URL.Query()["passM"][0]
		tokenU := r.URL.Query()["token"][0]

		if tokenU == token {
			if passMR == passM {
				var client Model.Client

				client.Name = nameR
				client.Surname = snameR
				client.Sn = snR
				client.RoomNumber = roomNumber
				client.CarID = carId

				json.NewDecoder(r.Body).Decode(&client)

				_ = db.QueryRow("insert into clients (name, sname, sn, room, carid) values($1, $2, $3, $4, $5);", client.Name, client.Surname, client.Sn, client.RoomNumber, client.CarID)
				http.Redirect(w, r, url+"logged.html", http.StatusSeeOther)
			} else {
				http.Redirect(w, r, url+"logged.html", http.StatusSeeOther)
			}
		} else {
			fmt.Fprintf(w, "Bad token verification")
		}
	}
}

func GetClients(db *sql.DB, aUrl string, token string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenU := r.URL.Query()["token"][0]
		if tokenU == token {
			var client Model.Client
			rows, _ := db.Query("select * from clients")
			defer rows.Close()
			i := 1
			for rows.Next() {
				rows.Scan(&client.Name, &client.Surname, &client.Sn, &client.RoomNumber, &client.CarID)
				is := strconv.Itoa(i)
				w.Header().Add("Content-Type", "html")
				fmt.Fprintf(w, View.ShowClients(client.Name, client.Surname, client.Sn, is, aUrl, tokenU))
				i++
			}
		} else {
			fmt.Fprintf(w, "Bad token verification")
		}
	}
}

func DeleteClient(db *sql.DB, url string, passMD string, token string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		snD := r.URL.Query()["sn"][0]
		passM := r.URL.Query()["passM"][0]
		tokenU := r.URL.Query()["token"][0]
		if token == tokenU {
			if passM == passMD {
				result, _ := db.Exec("delete from clients where sn=$1", snD)
				rowsUpdated, _ := result.RowsAffected()
				fmt.Println(rowsUpdated)
				http.Redirect(w, r, url+"logged.html", http.StatusSeeOther)
			} else {
				http.Redirect(w, r, url+"logged.html?error=wrong-manager-passD", http.StatusSeeOther)
			}
		} else {
			fmt.Fprintf(w, "Bad token verification")
		}
	}
}

func ShowClient(db *sql.DB, token string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenU := r.URL.Query()["token"][0]
		if token == tokenU {
			var client Model.Client
			sn := r.URL.Query()["sn"][0]
			rows := db.QueryRow("select * from clients where sn=$1", sn)
			rows.Scan(&client.Name, &client.Surname, &client.Sn, &client.RoomNumber, &client.CarID)
			fmt.Fprintf(w, View.ShowClient(client.Name, client.Surname, client.Sn, client.RoomNumber, client.CarID))
		} else {
			fmt.Fprintf(w, "Bad token verification")
		}
	}
}
