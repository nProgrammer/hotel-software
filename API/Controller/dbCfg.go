package Controller

import (
	"database/sql"
	"github.com/lib/pq"
	"log"
)

func DbCfg(db *sql.DB, dbUrl string) *sql.DB {
	pgUrl, _ := pq.ParseURL(dbUrl)
	log.Println(pgUrl)
	db, _ = sql.Open("postgres", pgUrl)
	_ = db.Ping()

	return db
}
