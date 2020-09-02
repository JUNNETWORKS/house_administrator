package handlers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// TODO:  GetControllerCommands ... 特定のコントローラーの全てのコマンドを取得
/*

レスポンス
{
	"status": "SUCCESS",
	"commands": [
		{
			"name": "turn on light",
			"command_id": 1
		}
	]
}
*/
func GetControllerCommands(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

}

// TODO: RegisterCommand ... コントローラーにコマンドを追加
/*
リクエスト
{
	"name": "turn off right"
}
レスポンス
{
	"command_id": 42
}
*/
func RegisterCommand(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

}
