package data

import (
	"log"
	"time"
)

type Room struct {
	Id          int
	Name        string
	Description string
	OwnerId     int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func Rooms() (rooms []Room, err error) {
	rows, err := Db.Query("SELECT id, name, description, owner_id, created_at, updated_at FROM Rooms")
	if err != nil {
		log.Print("Failed get all rooms")
		return
	}
	for rows.Next() {
		var room Room
		err = rows.Scan(room.Id, room.Name, room.Description, room.OwnerId, room.CreatedAt, room.UpdatedAt)
		if err != nil {
			return
		}
		rooms = append(rooms, room)
	}
	rows.Close()
	return
}
