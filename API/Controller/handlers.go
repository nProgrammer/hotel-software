package Controller

import (
	"API/Model"
	"API/Utils"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)


func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func Login(userL string, userP string, token string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		tokenU := r.URL.Query()["token"][0]
		var user Model.ClientLogin
		json.NewDecoder(r.Body).Decode(&user)
		fmt.Println(user)
		if user.Login == userL && user.Pass == userP && tokenU == token {
			json.NewEncoder(w).Encode(Utils.LoginU("SUCCESS"))
		} else {
			json.NewEncoder(w).Encode(Utils.LoginU("NOT SUCCESS"))
		}
	}
}

func RegisterClient(db *sql.DB, passM string, token string) http.HandlerFunc {
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
				json.NewEncoder(w).Encode(client)
			} else {
				json.NewEncoder(w).Encode("WRONG MANAGER PASSWORD")
			}
		} else {
			fmt.Fprintf(w, "Bad token verification")
		}
	}
}

func GetClients(db *sql.DB, token string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenU := r.URL.Query()["token"][0]
		if tokenU == token {
			var client Model.Client
			var clients []Model.Client
			rows, _ := db.Query("select * from clients")
			defer rows.Close()
			i := 1
			for rows.Next() {
				rows.Scan(&client.Name, &client.Surname, &client.Sn, &client.RoomNumber, &client.CarID)
				clients = append(clients, client)
				i++
			}
			json.NewEncoder(w).Encode(clients)
		} else {
			json.NewEncoder(w).Encode("BAD TOKEN VERIFICATION")
		}
	}
}

func DeleteClient(db *sql.DB, passMD string, token string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var clientJ Model.ClientDel
		json.NewDecoder(r.Body).Decode(&clientJ)
		tokenU := r.URL.Query()["token"][0]
		if token == tokenU {
			if clientJ.Passm == passMD {
				result, _ := db.Exec("delete from clients where sn=$1", clientJ.Sn)
				rowsUpdated, _ := result.RowsAffected()
				fmt.Println(rowsUpdated)
				json.NewEncoder(w).Encode(clientJ.Sn)
			} else {
				json.NewEncoder(w).Encode("BAD MANAGER PASSWORD")
			}
		} else {
			fmt.Fprintf(w, "BAD TOKEN VERIFICATION")
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
			json.NewEncoder(w).Encode(client)
		} else {
			json.NewEncoder(w).Encode("Bad token verification")
		}
	}
}
