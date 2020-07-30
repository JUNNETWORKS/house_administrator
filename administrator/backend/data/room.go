package data

import (
	"log"
	"time"
)

type Room struct {
	Id          int       `db:id`
	Name        string    `db:name`
	Description string    `db:description`
	OwnerId     int       `db:owner_id`
	CreatedAt   time.Time `db:created_at`
	UpdatedAt   time.Time `db:updated_at`
}

func Rooms() (rooms []Room, err error) {
	rows, err := Db.Queryx("SELECT * FROM rooms")
	if err != nil {
		// TODO: エラーログの出し方を考える
		log.Print("Failed get all rooms")
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
