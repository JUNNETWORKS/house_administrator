package data

import (
	"gorm.io/gorm"
)

// Room はDBにおけるroomsテーブルの各行を表す構造体
type Room struct {
	gorm.Model
	Name        string       `json:"name"`
	Description string       `json:"description"`
	HostName    string       `json:"host_name"`
	Sensors     []Sensor     `json:"sensors"`
	Controllers []Controller `json:"controllers"`
	// OwnerID     int      `db:"owner_id" json:"owner_id"`
}

// NewRoom  構造体Roomのポインタを返す
func NewRoom(name string, description string) *Room {
	room := &Room{Name: name, Description: description}
	return room
}

// GetRooms は全部屋のデータをDBから取得し,返す関数.
func GetRooms() ([]Room, error) {
	var rooms []Room
	result := Db.Find(&rooms)
	return rooms, result.Error
}

// RetrieveRoom はidのRoomを取得して返す.
func RetrieveRoom(id int) (*Room, error) {
	var room Room

	result := Db.First(&room, id)
	return &room, result.Error
}

// Create DBに挿入する
func (room *Room) Create() error {
	result := Db.Create(room)
	return result.Error
}

// Update Room構造体のIDのRoomをRoom構造体の情報を元に更新
func (room *Room) Update() error {
	result := Db.Save(room)
	return result.Error
}

// Delete Room構造体のIDのRoomをDBから削除
func (room *Room) Delete() error {
	result := Db.Delete(room)
	return result.Error
}

// DeleteByID ... IDのRoomをテーブルから削除
func DeleteByID(id int) error {
	result := Db.Delete(&Room{}, id)
	return result.Error
}

// DeleteRoomAll ... roomsテーブルの全てのレコードを削除する
// テストでのみの利用を想定している
func DeleteRoomAll() error {
	result := Db.Exec(`DELETE FROM rooms`)
	return result.Error
}
