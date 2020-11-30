package data

import (
	"gorm.io/gorm"
)

// Measurement ... 計測データを表す構造体
type Measurement struct {
	gorm.Model
	SensorID uint
	Value    float64
}

// Measurements ... 全計測データを取得する
func Measurements() ([]Measurement, error) {
	var measurements []Measurement
	result := Db.Find(&measurements)
	return measurements, result.Error
}
