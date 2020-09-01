package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/JUNNETWORKS/house_administrator/handlers"

	"github.com/julienschmidt/httprouter"
)

func init() {
	log.SetFlags(log.Ldate | log.Lshortfile)
}

func dummyHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "[%s] Response for you :)", time.Now().Format(time.RFC3339Nano))
}

func main() {

	// マルチプレクサにはhttprouterを使う
	mux := httprouter.New()

	mux.GET("/", dummyHandler)

	// 部屋一覧
	mux.GET("/rooms", handlers.GetRooms)
	// 部屋を追加
	mux.POST("/rooms", handlers.RegisterRoom)

	// 特定の部屋の操作
	mux.GET("/rooms/:roomID", handlers.GetRoom)       // 部屋のトップページ
	mux.PUT("/rooms/:roomID", handlers.UpdateRoom)    // 部屋の情報を更新
	mux.DELETE("/rooms/:roomID", handlers.DeleteRoom) // 部屋を削除

	// センサー関連
	mux.GET("/rooms/:roomID/sensors", handlers.GetSensors)      // 搭載されている全センサー
	mux.POST("/rooms/:roomID/sensors", handlers.RegisterSensor) // センサーを追加

	// 特定のセンサーの操作
	mux.GET("/rooms/:roomID/sensors/:sensorID", handlers.GetSensor)       // 特定のセンサーに関する情報
	mux.PUT("/rooms/:roomID/sensors/:sensorID", handlers.UpdateSensor)    // 特定のセンサーを更新
	mux.DELETE("/rooms/:roomID/sensors/:sensorID", handlers.DeleteSensor) // 特定のセンサーを削除

	// 特定のセンサーの記録
	mux.GET("/rooms/:roomID/sensors/:sensorID/measurements", handlers.GetMeasurements)      // 特定のセンサーの記録を返す
	mux.POST("/rooms/:roomID/sensors/:sensorID/measurements", handlers.RegisterMeasurement) // 特定のセンサーの記録を記録
	mux.GET("/rooms/:roomID/sensors/:sensorID/latest", handlers.GetLatestMeasurement)       // 特定のセンサーの最新の記録を返す

	// コントローラー関係
	mux.GET("/rooms/:roomID/controllers", handlers.GetAllControllersInRoom) // 利用可能な全コントローラー
	mux.POST("/rooms/:roomID/controllers", handlers.RegisterController)     // コントローラーを追加

	// 特定のコントローラーの情報を取得
	mux.GET("/rooms/:roomID/controllers/:controllerID", handlers.GetController) // コントローラーの情報や利用可能なコマンドを返す

	// 特定のコントローラーのコマンド
	mux.GET("/rooms/:roomID/controllers/:controllerID/commands", dummyHandler)  // コマンド一覧
	mux.POST("/rooms/:roomID/controllers/:controllerID/commands", dummyHandler) // コマンドを追加

	// 特定のコントローラーの特定のコマンド
	mux.GET("/rooms/:roomID/controllers/:controllerID/commands/:commandID", dummyHandler)    // コマンドの情報
	mux.POST("/rooms/:roomID/controllers/:controllerID/commands/:commandID", dummyHandler)   // コマンド実行
	mux.PUT("/rooms/:roomID/controllers/:controllerID/commands/:commandID", dummyHandler)    // コマンドを更新
	mux.DELETE("/rooms/:roomID/controllers/:controllerID/commands/:commandID", dummyHandler) // コマンドを削除

	server := http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	log.Fatal(server.ListenAndServe())
}
