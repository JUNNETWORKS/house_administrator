package data

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// sqlにアクセスするためのパッケージ変数
var Db *sqlx.DB

func init() {
	var err error
	// TODO: DBの名前は変えるかも
	Db, err = sqlx.Open("postgres", "dbname=house sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	return
}
