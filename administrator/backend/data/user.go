package data

import "gorm.io/gorm"

// User ... ユーザーを表す構造体
type User struct {
	gorm.Model
	Name  string
	Rooms []Room `gorm:"many2many:user_rooms;"`
}

// Users ... 全ユーザーを取得する関数
func Users() ([]User, error) {
	var users []User
	result := Db.Find(&users)
	return users, result.Error
}
