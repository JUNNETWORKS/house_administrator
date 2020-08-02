package data

import "time"

// Sensor ... センサー情報を表す構造体
type Sensor struct {
	ID           int
	SensorTypeID int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// SensorType ... センサーの種類を表す構造体
type SensorType struct {
	ID   int
	Name string
	Unit string
}

// Sensors ... 全センサーを取得する
func Sensors() ([]Sensor, error) {
	var sensors []Sensor
	err := Db.Select(&sensors, "SELECT * FROM sensors")
	return sensors, err
}
