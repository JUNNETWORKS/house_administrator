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
	rows, err := Db.Queryx("SELECT * FROM rooms")
	if err != nil {
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
