package data

import "time"

type User struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func Users() (users []User, err error) {
	err = Db.Select(&users, "SELECT * FROM users")
	return
}
