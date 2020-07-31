package data

import "time"

type Sensor struct {
	Id           int
	SensorTypeId int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type SensorType struct {
	Id   int
	Name string
	Unit string
}

func Sensors() (sensors []Sensor, err error) {
	err = Db.Select(&sensors, "SELECT * FROM sensors")
	return
}
