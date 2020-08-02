package data

import "time"

// User ... ユーザーを表す構造体
type User struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Users ... 全ユーザーを取得する関数
func Users() ([]User, error) {
	var users []User
	err := Db.Select(&users, "SELECT * FROM users")
	return users, err
}
