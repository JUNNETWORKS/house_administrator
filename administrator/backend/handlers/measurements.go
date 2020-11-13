package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// TODO: GetMeasurements ... センサーIDで指定されたセンサーで取得したデータを指定された期間分取得する
/*
ex:

リクエスト
GET /rooms/1/sensors/25/measurements?from=2020-08-12T00:00:00&until=2020-08-14T00:00:00

レスポンス
{
	"status": "SUCCESS",
	"message": "SUCCESS",
	"measurements": [  // statusがFAILEDの時は空のリスト
		{
			"datetime": "2020-08-12T00:00:00",
			"measurement": 20.8
		},
		...
	]
}
*/
func GetMeasurements(c *gin.Context) {
}

// TODO: GetLatestMeasurement ... 特定のセンサーの最新の記録を返す
/*
ex:

リクエスト
GET /rooms/1/sensors/25/measurements/latest

レスポンス
{
	"status": "SUCCESS",
	"message": "SUCCESS",
	"measurements": {  // statusがFAILEDの時は空のオブジェクト
			"datetime": "2020-08-12T00:00:00",
			"measurement": 20.8
		}
}
*/
func GetLatestMeasurement(c *gin.Context) {
}

// TODO: RegisterMeasurement ... 特定のセンサーの記録を記録
/*
ex:

リクエスト
POST /rooms/1/sensors/25/measurements
{
	"datetime": "2020-09-12T01:05:00",  // ない場合はサーバー側で現在時刻を設定
	"measurement": 20.4
}

レスポンス
成功したらステータスコード200を返す.
エラーの場合は相応のステータスコードとボディにエラーメッセージ
*/
func RegisterMeasurement(c *gin.Context) {

}
