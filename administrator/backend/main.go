package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

func main() {

	// マルチプレクサにはhttprouterを使う
	mux := httprouter.New()

	// 利用可能な部屋を返す
	mux.GET("/rooms", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprintf(w, "[%s] Response all rooms", time.Now().Format(time.RFC3339Nano))
	})
	// 部屋のトップページを返す
	mux.GET("/rooms/:room", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		roomName := p.ByName("room")
		fmt.Fprintf(w, "[%s] This room name is %s", time.Now().Format(time.RFC3339Nano), roomName)
	})

	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}
	server.ListenAndServe()

	// TODO: ここから下も修正する

	// // 利用可能な操作を返す
	// r.GET("/rooms/:room/controllers", func(c *gin.Context) {})
	// // エアコンの操作に関する情報を返す (操作出来る項目や値の範囲など)
	// r.GET("/rooms/:room/controllers/air-conditioner", func(c *gin.Context) {
	// 	roomName := c.Param("room")
	// 	fmt.Printf("%s", roomName)
	// })
	// // エアコンに対して何かしらの操作を送信するAPI
	// r.POST("/rooms/:room/controllers/air-conditioner", func(c *gin.Context) {})
	// // 将来的には以下の行の形式に置き換わる予定
	// r.GET("/rooms/:room/controllers/*device", func(c *gin.Context) {})
	// r.POST("/rooms/:room/controllers/*device", func(c *gin.Context) {})
	// // 利用可能な計測値の種類を返す (温度や湿度,CO2濃度など)
	// r.GET("/rooms/:room/measurements", func(c *gin.Context) {})
	// // 部屋の温度の詳細やグラフを表示するページを返す
	// r.GET("/rooms/:room/measurements/temp", func(c *gin.Context) {})
}
