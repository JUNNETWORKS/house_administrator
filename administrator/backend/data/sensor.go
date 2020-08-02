package data

import "time"

type Sensor struct {
	ID           int
	SensorTypeID int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type SensorType struct {
	ID   int
	Name string
	Unit string
}

func Sensors() ([]Sensor, error) {
	var sensors []Sensor
	err := Db.Select(&sensors, "SELECT * FROM sensors")
	return sensors, err
}
