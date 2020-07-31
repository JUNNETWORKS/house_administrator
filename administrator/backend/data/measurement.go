package data

import (
	"github.com/jmoiron/sqlx"
	"log"
	"time"
)

type Measurement struct {
	Id        int
	SensorId  int
	Value     float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func Measurements() (measurements []Measurements, err error) {
	rows, err := Db.Queryx("SELECT * FROM measurements")
	if err != nil {
		log.Println("Failed to get all measurements")
		return
	}
	for rows.Next() {
		var measurement Measurement
		err = rows.StructScan(&measurement)
		if err != nil {
			return
		}
		measurements = append(measurements, measurement)
	}
	return
}
