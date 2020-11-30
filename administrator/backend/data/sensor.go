package data

import (
	"gorm.io/gorm"
)

// Sensor センサー情報を表す構造体
type Sensor struct {
	gorm.Model
	RoomID       uint
	SensorTypeID uint
}

// SensorType  センサーの種類を表す構造体
type SensorType struct {
	gorm.Model
	Name string
	Unit string
}

// Sensors  全センサーを取得する
func Sensors() ([]Sensor, error) {
	var sensors []Sensor
	result := Db.Find(&sensors)
	return sensors, result.Error
}

// GetSensorsByID  部屋のID紐付いているセンサーを全て取得する
func GetSensorsByID(roomID int) ([]Sensor, error) {
	var sensors []Sensor
	result := Db.Where("room_id = ?", roomID).Find(sensors)
	return sensors, result.Error
}
