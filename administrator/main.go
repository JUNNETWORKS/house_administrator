package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	grpc "./grpc"
)

func main() {
	grpc.Serve()

	r := gin.Default()

	// 利用可能な部屋を返す
	r.POST("/rooms", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"api_list": []string{
				"jun_room",
				"garage",
			},
		})
	})
	// 部屋のトップページを返す
	r.GET("/rooms/:room/", func(c *gin.Context) {})
	// 利用可能な操作を返す
	r.GET("/rooms/:room/controllers", func(c *gin.Context) {})
	// エアコンの操作に関する情報を返す (操作出来る項目や値の範囲など)
	r.GET("/rooms/:room/controllers/air-conditioner", func(c *gin.Context) {
		roomName := c.Param("room")
		fmt.Printf("%s", roomName)
	})
	// エアコンに対して何かしらの操作を送信するAPI
	r.POST("/rooms/:room/controllers/air-conditioner", func(c *gin.Context) {})
	// 将来的には以下の行の形式に置き換わる予定
	r.GET("/rooms/:room/controllers/*device", func(c *gin.Context) {})
	r.POST("/rooms/:room/controllers/*device", func(c *gin.Context) {})
	// 利用可能な計測値の種類を返す (温度や湿度,CO2濃度など)
	r.GET("/rooms/:room/measurements", func(c *gin.Context) {})
	// 部屋の温度の詳細やグラフを表示するページを返す
	r.GET("/rooms/:room/measurements/temp", func(c *gin.Context) {})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
