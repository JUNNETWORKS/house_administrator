package data

import (
	"time"
)

// Command ... コントローラーが利用できるコマンド
type Command struct {
	ControllerID uint
	Controller   Controller `json:"controller_id"`
	Name         string     `json:"name"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}
