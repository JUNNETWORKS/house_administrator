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

func Sensors() (sensors []Sensor, err error) {
	err = Db.Select(&sensors, "SELECT * FROM sensors")
	return
}
