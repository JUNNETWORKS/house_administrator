package data

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

// sqlにアクセスするためのパッケージ変数
var Db *sql.DB

func init() {
	var err error
	// TODO: DBの名前は変えるかも
	Db, err = sql.Open("postgres", "dbname=house sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	return
}
