package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/JUNNETWORKS/house_administrator/data"
	"github.com/gin-gonic/gin"
)

// Binding from JSON
type RequestRoomBody struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"-"`
}

// GetRooms ... 全ての部屋の情報をJSONで返す
// レスポンスの例
// [{'id': 1,
//   'name': 'JUN Room',
//   'description': 'this is test room',
//   'owner_id': 1,
//   'created_at': '2000-01-01T00:00:00Z',
//   'updated_at': '2000-01-01T00:00:00Z'}]
func GetRooms(c *gin.Context) {
	rooms, err := data.GetRooms()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed get all rooms"})
	}
	log.Println("全部屋の情報を取得")
	c.JSON(http.StatusOK, rooms)
}

// RegisterRoom ... 新しい部屋を登録する
// リクエストの例
// {
// 	"name": "New Room",
// 	"description": "This is new room"
// 	"owner_id": 1
// }
func RegisterRoom(c *gin.Context) {
	var err error
	var newRoom *RequestRoomBody
	if err = c.BindJSON(newRoom); err != nil {
		log.Println("リクエストが間違っています.")
		log.Println("リクエストボディから構造体にマッピング出来ませんでした.")
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	room := data.NewRoom(newRoom.Name, newRoom.Description)
	err = room.Create()
	if err != nil {
		log.Println("新しい部屋をDBに挿入出来ませんでした.")
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	log.Printf("新しい部屋を追加しました\n%v\n", room)
	log.Printf("RoomID %d の部屋を作成\n", room.ID)
	c.JSON(http.StatusOK, room)
}

// GetRoom ... 部屋の情報を取得して返す
func GetRoom(c *gin.Context) {
	roomID, err := strconv.Atoi(c.Param("roomID"))
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	room, err := data.RetrieveRoom(roomID)
	if err != nil {
		log.Printf("RoomID %d のデータを取得できませんでした.\n", roomID)
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	log.Printf("RoomID %d の部屋を取得\n", room.ID)
	c.JSON(http.StatusOK, room)
}

// UpdateRoom ... 部屋の情報を更新する
// 例: descriptionを変更する
// {
// 	"description": "let's change description!"
// }
func UpdateRoom(c *gin.Context) {
	roomID, err := strconv.Atoi(c.Param("roomID"))
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// 現在の部屋のデータを取得
	room, err := data.RetrieveRoom(roomID)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	// リクエストボディをパースしてJSONを取り出す
	len := c.Request.ContentLength
	body := make([]byte, len)
	_, err = c.Request.Body.Read(body)
	if err != nil && err.Error() != "EOF" {
		log.Printf("%s", string(body))
		log.Println("リクエストが間違っています.")
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	// JSONで指定されたデータroomを更新する
	err = json.Unmarshal(body, &room)
	if err != nil {
		log.Println("リクエストボディから構造体にマッピング出来ませんでした.")
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	// DBに変更を反映する
	if err = room.Update(); err != nil {
		log.Println("部屋のデータを更新出来ませんでした.")
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	log.Printf("RoomID %d の部屋を更新\n", room.ID)
	c.JSON(http.StatusOK, room)
}

// DeleteRoom ... 部屋を削除する
func DeleteRoom(c *gin.Context) {
	var (
		roomID int
		room   *data.Room
		err    error
	)
	if roomID, err = strconv.Atoi(c.Param("roomID")); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// 現在の部屋のデータを取得
	if room, err = data.RetrieveRoom(roomID); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	if err = room.Delete(); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	log.Printf("RoomID %d の部屋を削除\n", room.ID)
	c.Status(http.StatusOK)
}
