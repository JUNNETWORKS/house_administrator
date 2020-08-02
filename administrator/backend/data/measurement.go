package data

import (
	"log"
	"time"
)

// Measurement ... 計測データを表す構造体
type Measurement struct {
	ID        int
	SensorID  int
	Value     float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Measurements ... 全計測データを取得する
func Measurements() ([]Measurement, error) {
	rows, err := Db.Queryx("SELECT * FROM measurements")
	if err != nil {
		log.Println("Failed to get all measurements")
		return nil, err
	}
	var measurements []Measurement
	for rows.Next() {
		var measurement Measurement
		err = rows.StructScan(&measurement)
		if err != nil {
			return nil, err
		}
		measurements = append(measurements, measurement)
	}
	return measurements, nil
}
