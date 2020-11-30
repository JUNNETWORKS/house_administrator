package data

import (
	"gorm.io/gorm"
)

// Controller コントローラー情報を表す構造体
type Controller struct {
	gorm.Model
	Name   string `json:"name"`
	RoomID uint   `json:"-"`
}

// GetRoomControllers 部屋のIDを元にその部屋に紐付いているControllerを返す
func GetRoomControllers(roomID int) ([]Controller, error) {
	var controllers []Controller
	result := Db.Where("room_id = ?", roomID).Find(&controllers)
	return controllers, result.Error
}

// Create DBに挿入する
func (controller *Controller) Create() error {
	result := Db.Create(controller)
	return result.Error
}
