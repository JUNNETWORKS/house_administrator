package handlers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// TODO: GetSensors ... 部屋にある全てのセンサーを取得し, JSONで返す
func GetSensors(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// roomID, err := strconv.Atoi(p.ByName("roomID"))

}

// TODO: RegisterSensor ... 部屋に新しいセンサーを登録する
func RegisterSensor(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}

// TODO: GetSensor ... 特定のセンサーの情報を取得
func GetSensor(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}

// TODO: UpdateSensor ... 特定のセンサーの情報を更新
func UpdateSensor(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}

// TODO: DeleteSensor ... 特定のセンサーを削除
func DeleteSensor(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}
