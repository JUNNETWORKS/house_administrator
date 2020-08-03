package data

import (
	"fmt"
	"log"

	"github.com/JUNNETWORKS/house_administrator/utils"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// Db はsqlにアクセスするためのパッケージ変数
var Db *sqlx.DB

func init() {
	var err error
	// TODO: DBの名前は変えるかも
	dbhost := utils.GetEnvWithDefault("DB_HOST", "db")
	dbport := utils.GetEnvWithDefault("DB_PORT", "5432")
	dbname := utils.GetEnvWithDefault("DB_NAME", "house")
	dbuser := utils.GetEnvWithDefault("DB_USER", "postgres")
	dbpass := utils.GetEnvWithDefault("DB_PASS", "postgres")
	dataSourceName := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable", dbhost, dbport, dbname, dbuser, dbpass)
	Db, err = sqlx.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	return
}
