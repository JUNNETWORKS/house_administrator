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
