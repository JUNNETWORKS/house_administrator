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
}

// RegisterRoom ... 新しい部屋を登録する
// リクエストの例
// {
// 	"name": "New Room",
// 	"description": "This is new room"
// 	"owner_id": 1
// }
func RegisterRoom(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	len := r.ContentLength
	body := make([]byte, len)
	_, err := r.Body.Read(body)
	if err != nil && err.Error() != "EOF" {
		log.Printf("%s", string(body))
		log.Println("リクエストが間違っています.")
		log.Println(err.Error())
		w.Write([]byte(err.Error()))
		return
	}
	var room *data.Room
	err = json.Unmarshal(body, &room)
	if err != nil {
		log.Println("リクエストボディから構造体にマッピング出来ませんでした.")
		log.Println(err.Error())
		w.Write([]byte(err.Error()))
		return
	}
	err = room.Create()
	if err != nil {
		log.Println("新しい部屋をDBに挿入出来ませんでした.")
		log.Println(err.Error())
		w.Write([]byte(err.Error()))
		return
	}
	log.Printf("新しい部屋を追加しました\n%v\n", room)
	resJSON, err := json.MarshalIndent(room, "", "\t")
	if err != nil {
		log.Println("新しい部屋についてのJSONを作成出来ませんでした.")
		log.Println(err.Error())
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(resJSON)
}
