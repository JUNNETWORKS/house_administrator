package data

import "time"

type Account struct {
	Id        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
