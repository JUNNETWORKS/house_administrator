package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
func GetControllerCommands(c *gin.Context) {

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
func RegisterCommand(c *gin.Context) {

}
