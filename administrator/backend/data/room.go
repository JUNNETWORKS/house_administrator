package data

import (
	"log"
	"time"
)

// Room はDBにおけるroomsテーブルの各行を表す構造体
type Room struct {
	ID          int       `db:"id" json:"id"`
	Name        string    `db:"name" json:"name"`
	Description string    `db:"description" json:"description"`
	HostName    string    `db:"host" json:"host_name"`
	OwnerID     int       `db:"owner_id" json:"owner_id"`
	Sensors     []Sensor  `json:"sensors"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}

// NewRoom ... 構造体Roomのポインタを返す
func NewRoom(name string, description string) *Room {
	room := &Room{Name: name, Description: description}
	return room
}

// CreatedAtDate はCreatedAt属性を文字列にして返す
func (room *Room) CreatedAtDate() string {
	return room.CreatedAt.Format(time.RFC3339)
}

// UpdatedAtDate はUpdatedAt属性を文字列にして返す
func (room *Room) UpdatedAtDate() string {
	return room.UpdatedAt.Format(time.RFC3339)
}

// GetRooms は全部屋のデータをDBから取得し,返す関数.
func GetRooms() ([]Room, error) {
	var rooms []Room
	rows, err := Db.Queryx("SELECT * FROM rooms")
	if err != nil {
		// TODO: エラーログの出し方を考える
		log.Println("Failed get all rooms")
		return nil, err
	}
	for rows.Next() {
		var room Room
		err = rows.StructScan(&room)
		if err != nil {
			return nil, err
		}
		rooms = append(rooms, room)
	}
	rows.Close()
	return rooms, nil
}

// RetrieveRoom はidのRoomを取得して返す.
func RetrieveRoom(id int) (*Room, error) {
	var room Room
	err := Db.Get(&room, "SELECT * FROM rooms WHERE id = $1", id)
	return &room, err
}

// Create DBに挿入する
func (room *Room) Create() error {
	now := time.Now()
	room.CreatedAt = now
	room.UpdatedAt = now
	statement := `
	INSERT INTO rooms
    (name, description, owner_id, created_at, updated_at)
	VALUES(:name, :description, :owner_id, :created_at, :updated_at)
	RETURNING id;
	`
	stmt, err := Db.PrepareNamed(statement)
	if err != nil {
		return err
	}
	defer stmt.Close()
	err = stmt.QueryRow(map[string]interface{}{
		"name":        room.Name,
		"description": room.Description,
		"owner_id":    room.OwnerID,
		"created_at":  room.CreatedAtDate(),
		"updated_at":  room.UpdatedAtDate(),
	}).StructScan(room)
	return nil
}

// Update Room構造体のIDのRoomをRoom構造体の情報を元に更新
func (room *Room) Update() error {
	now := time.Now()
	room.UpdatedAt = now
	schema := `
	UPDATE rooms
	SET name = $1
	, description = $2
	, owner_id = $3
	, updated_at= $4
	WHERE id = $5;
	`
	_, err := Db.Exec(schema, room.Name, room.Description, room.OwnerID, room.UpdatedAtDate(), room.ID)
	return err
}

// Delete Room構造体のIDのRoomをDBから削除
func (room *Room) Delete() error {
	schema := `
	DELETE FROM rooms
	WHERE id = $1;
	`
	_, err := Db.Exec(schema, room.ID)
	return err
}

// DeleteByID ... IDのRoomをテーブルから削除
func DeleteByID(id int) error {
	schema := `
	DELETE FROM rooms
	WHERE id = $1;
	`
	_, err := Db.Exec(schema, id)
	return err
}

// DeleteRoomAll ... roomsテーブルの全てのレコードを削除する
// テストでのみの利用を想定している
func DeleteRoomAll() error {
	schema := `
	DELETE FROM rooms;
	`
	_, err := Db.Exec(schema)
	return err
}
