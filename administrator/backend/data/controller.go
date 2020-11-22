package data

import (
	"time"
)

// Controller ... センサー情報を表す構造体
type Controller struct {
	ID        int       `db:"id"`
	Room      *Room     `db:"room_id" json:"-"`
	CreatedAt time.Time `db:"create_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
