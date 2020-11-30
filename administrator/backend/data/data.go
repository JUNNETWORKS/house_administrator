package data

import (
	"fmt"
	"log"

	"github.com/JUNNETWORKS/house_administrator/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Db はsqlにアクセスするためのパッケージ変数
var Db *gorm.DB

func init() {
	var err error
	// TODO: DBの名前は変えるかも
	dbhost := utils.GetEnvWithDefault("DB_HOST", "db")
	dbport := utils.GetEnvWithDefault("DB_PORT", "5432")
	dbname := utils.GetEnvWithDefault("DB_NAME", "house")
	dbuser := utils.GetEnvWithDefault("DB_USER", "postgres")
	dbpass := utils.GetEnvWithDefault("DB_PASS", "postgres")
	dataSourceName := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable", dbhost, dbport, dbname, dbuser, dbpass)
	Db, err = gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Auto Migarations
	Db.AutoMigrate(&User{}, &Room{}, &Controller{}, &Command{}, &Sensor{}, &Measurement{})
	return
}
