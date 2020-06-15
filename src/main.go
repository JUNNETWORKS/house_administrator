package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	/*
		フロントエンド
	*/
	// htmlのディレクトリを指定
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"message": "This is index page",
		})
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	/*
		バックエンド
	*/
	// 利用可能な部屋を返す
	r.POST("/api", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"api_list": []string{
				"jun_room",
				"garage",
			},
		})
	})
	// 利用可能な操作を返す (エアコンや照明操作)
	r.GET("/api/:room/", func(c *gin.Context) {})
	// エアコンに対して何かしらの操作を送信するAPI
	r.POST("/api/:room/air-conditioner", func(c *gin.Context) {
		roomName := c.Param("room")
		fmt.Printf("%s", roomName)
	})
	// 利用可能な値の種類を返す (温度や湿度,CO2濃度など)
	r.GET("/api/:room/metrics", func(c *gin.Context) {})
	// 部屋の温度を取得するAPI
	r.GET("/api/:room/metrics/:", func(c *gin.Context) {})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
