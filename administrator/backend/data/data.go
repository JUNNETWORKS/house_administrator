package data

import (
	"fmt"
	"github.com/JUNNETWORKS/house_administrator/utils"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// sqlにアクセスするためのパッケージ変数
var Db *sqlx.DB

func init() {
	var err error
	// TODO: DBの名前は変えるかも
	var dbhost, dbport, dbname string
	dbhost = utils.GetEnvWithDefault("DB_HOST", "db")
	dbport = utils.GetEnvWithDefault("DB_PORT", "5432")
	dbname = utils.GetEnvWithDefault("DB_NAME", "house")
	dataSourceName := fmt.Sprintf("host=%s port=%s dbname=%s sslmode=disable", dbhost, dbport, dbname)
	Db, err = sqlx.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	return
}
