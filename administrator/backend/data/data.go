package data

import (
	"github.com/JUNNETWORKS/house_administrator/utils"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// sqlにアクセスするためのパッケージ変数
var Db *sqlx.DB

func init() {
	var err error
	// TODO: DBの名前は変えるかも
	if dbport, ok := os.LookupEnv("DB_PORT"); !ok {
		dbport = "5432"
	}
	if dbname, ok := os.LookupEnv("DB_NAME"); !ok {
		dbname = "house"
	}
	dataSourceName := fmt.Sprintf("port=%s dbname=%s sslmode=disable", dbport, dbname)
	Db, err = sqlx.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	return
}
