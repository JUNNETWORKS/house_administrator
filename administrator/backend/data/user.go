package data

import "time"

type User struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func Users() ([]User, error) {
	var users []User
	err := Db.Select(&users, "SELECT * FROM users")
	return users, err
}
