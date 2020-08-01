package data

import (
	"log"
	"time"
)

// Room はDBにおけるroomsテーブルの各行を表す構造体
type Room struct {
	ID          int       `db:"id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	OwnerID     int       `db:"owner_id"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

// CreatedAtDate はCreatedAt属性を文字列にして返す
func (room *Room) CreatedAtDate() string {
	return room.CreatedAt.Format(time.RFC3339)
}

// UpdatedAtDate はUpdatedAt属性を文字列にして返す
func (room *Room) UpdatedAtDate() string {
	return room.UpdatedAt.Format(time.RFC3339)
}

// Rooms は全部屋のデータをDBから取得し,返す関数.
func Rooms() (rooms []Room, err error) {
	rows, err := Db.Queryx("SELECT * FROM rooms")
	if err != nil {
		// TODO: エラーログの出し方を考える
		log.Println("Failed get all rooms")
		return
	}
	for rows.Next() {
		var room Room
		err = rows.StructScan(&room)
		if err != nil {
			return
		}
		rooms = append(rooms, room)
	}
	rows.Close()
	return
}

// retrieveRoom はidのRoomを取得して返す.
func retrieveRoom(id int) (room *Room, err error) {
	err = Db.Select(&room, "SELECT * FROM rooms WHERE id = ?", id)
	return
}

// Create DBに挿入する
func (room *Room) Create() (err error) {
	now := time.Now()
	room.CreatedAt = now
	room.UpdatedAt = now
	schema := `
	INSERT INTO rooms
    (name, description, owner_id, created_at, updated_at)
	VALUES(?, ?, ?, ?, ?);
	`
	_, err = Db.Exec(schema, room.Name, room.Description, room.OwnerID, room.CreatedAtDate(), room.UpdatedAtDate())
	return
}

// Update Room構造体の情報を元に
func (room *Room) Update() (err error) {
	now := time.Now()
	room.UpdatedAt = now
	schema := `
	UPDATE rooms
	SET name = ?
	, description = ?
	, owner_id = ?
	, updated_at= ?
	WHERE id = ?;
	`
	_, err = Db.Exec(schema, room.Name, room.Description, room.OwnerID, room.UpdatedAtDate(), room.ID)
	return
}

// Delete Room構造体を元に
func (room *Room) Delete() (err error) {
	schema := `
	DELETE FROM rooms
	WHERE id = ?;
	`
	_, err = Db.Exec(schema, room.ID)
	return err
}
