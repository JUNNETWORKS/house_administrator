package data

import (
	"log"
	"time"
)

// Sensor ... センサー情報を表す構造体
type Sensor struct {
	ID         int   `db:"id"`
	Room       *Room `db:"room_id" json:"-"`
	SensorType *SensorType
	CreatedAt  time.Time `db:"create_at"`
	UpdatedAt  time.Time `db:"updated_at"`
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

// GetSensorsByID ... 部屋のID紐付いているセンサーを全て取得する
func GetSensorsByID(roomID int) ([]Sensor, error) {
	var sensors []Sensor
	rows, err := Db.Queryx("SELECT * FROM sensors WHERE room_id = $1", roomID)
	if err != nil {
		// TODO: エラーログの出し方を考える
		log.Printf("Failed get room No %d all sensors", roomID)
		return nil, err
	}
	for rows.Next() {
		var sensor Sensor
		err = rows.StructScan(&sensor)
		if err != nil {
			return nil, err
		}
		sensor.SensorType, err = GetSensorTypeByID(sensor.ID)
		if err != nil {
			return nil, err
		}
		sensors = append(sensors, sensor)
	}
	rows.Close()
	return sensors, nil
}

// GetSensorTypeByID ... SensorTypeID を元に SensorType{} を返す
func GetSensorTypeByID(SensorTypeID int) (*SensorType, error) {
	var sensorType SensorType
	err := Db.Select(&sensorType, "SELECT * FROM sensor_types WHERE $1", SensorTypeID)
	return &sensorType, err
}
