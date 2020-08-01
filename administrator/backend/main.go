package main

import (
	"fmt"
	"github.com/JUNNETWORKS/house_administrator/data"
	"log"
	"net/http"
	"time"

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
	mux.GET("/rooms", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		rooms, err := data.Rooms()
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		for i, room := range rooms {
			str := fmt.Sprintln("%d: %v", i, room)
			w.Write([]byte(str))
		}
	})
	// 利用可能な部屋
	mux.POST("/rooms", dummyHandler) // 部屋を追加

	// 特定の部屋の操作
	mux.GET("/rooms/:roomId", dummyHandler)    // 部屋のトップページ
	mux.PUT("/rooms/:roomId", dummyHandler)    // 部屋の情報を更新
	mux.DELETE("/rooms/:roomId", dummyHandler) // 部屋を削除

	// センサー関連
	mux.GET("/rooms/:roomId/sensors", dummyHandler)  // 搭載されている全センサー
	mux.POST("/rooms/:roomId/sensors", dummyHandler) // センサーを追加

	// 特定のセンサーの操作
	mux.GET("/rooms/:roomId/sensors/:sensorId", dummyHandler)    // 特定のセンサーに関する情報
	mux.PUT("/rooms/:roomId/sensors/:sensorId", dummyHandler)    // 特定のセンサーを更新
	mux.DELETE("/rooms/:roomId/sensors/:sensorId", dummyHandler) // 特定のセンサーを削除

	// 特定のセンサーの記録
	mux.GET("/rooms/:roomId/sensors/:sensorId/measurements", dummyHandler)  // 特定のセンサーの全記録を返す
	mux.POST("/rooms/:roomId/sensors/:sensorId/measurements", dummyHandler) // 特定のセンサーの記録を記録

	// コントローラー関係
	mux.GET("/rooms/:roomId/controllers", dummyHandler)  // 利用可能な全コントローラー
	mux.POST("/rooms/:roomId/controllers", dummyHandler) // コントローラーを追加

	// 特定のコントローラーの操作
	mux.GET("/rooms/:roomId/controllers/:controllerId", dummyHandler) // コントローラーの情報

	// 特定のコントローラーのコマンド
	mux.GET("/rooms/:roomId/controllers/:controllerId/commands", dummyHandler)  // コマンド一覧
	mux.POST("/rooms/:roomId/controllers/:controllerId/commands", dummyHandler) // コマンドを追加

	// 特定のコントローラーの特定のコマンド
	mux.GET("/rooms/:roomId/controllers/:controllerId/commands/:commandId", dummyHandler)    // コマンドの情報
	mux.POST("/rooms/:roomId/controllers/:controllerId/commands/:commandId", dummyHandler)   // コマンド実行
	mux.PUT("/rooms/:roomId/controllers/:controllerId/commands/:commandId", dummyHandler)    // コマンドを更新
	mux.DELETE("/rooms/:roomId/controllers/:controllerId/commands/:commandId", dummyHandler) // コマンドを削除

	server := http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
