package handlers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// TODO: GetAllControllersInRoom ... 部屋の中にある全てのコントローラーを取得
func GetAllControllersInRoom(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
}

// TODO: RegisterController ... 部屋に新しいコントローラーを追加
func RegisterController(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

}

// TODO: GetController ... 特定のコントローラーの情報や利用可能なコマンドを返す
func GetController(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

}
