package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/JUNNETWORKS/house_administrator/data"
	"github.com/julienschmidt/httprouter"
)

// GetRooms ... 全ての部屋の情報をJSONで返す
// レスポンスの例
// [{'id': 1,
//   'name': 'JUN Room',
//   'description': 'this is test room',
//   'owner_id': 1,
//   'created_at': '2000-01-01T00:00:00Z',
//   'updated_at': '2000-01-01T00:00:00Z'}]
func GetRooms(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	rooms, err := data.GetRooms()
	if err != nil {
		log.Println(err)
	}
	output, err := json.Marshal(&rooms)
	if err != nil {
		log.Println(err)
	}
	w.Write(output)
	return
}

func AddRoom(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	return
}
