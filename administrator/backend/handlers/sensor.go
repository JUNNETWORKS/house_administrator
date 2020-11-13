package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/julienschmidt/httprouter"
)

// TODO: GetSensors ... 部屋にある全てのセンサーを取得し, JSONで返す
func GetSensors(c *gin.Context) {
	// roomID, err := strconv.Atoi(p.ByName("roomID"))

}

// TODO: RegisterSensor ... 部屋に新しいセンサーを登録する
func RegisterSensor(c *gin.Context) {

}

// TODO: GetSensor ... 特定のセンサーの情報を取得
func GetSensor(c *gin.Context) {

}

// TODO: UpdateSensor ... 特定のセンサーの情報を更新
func UpdateSensor(c *gin.Context) {

}

// TODO: DeleteSensor ... 特定のセンサーを削除
func DeleteSensor(c *gin.Context) {

}
