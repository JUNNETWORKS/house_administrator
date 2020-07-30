package data

import "time"

type Measurement struct {
	Id        int
	SensorId  int
	Value     float64
	CreatedAt time.Time
	UpdatedAt time.Time
}
