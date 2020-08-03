package handlers

import (
	"github.com/JUNNETWORKS/house_administrator/data"
)

// テスト用に全テーブルのレコード削除を行う
func setup() {
	data.DeleteRoomAll()
}
