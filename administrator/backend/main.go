package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/JUNNETWORKS/house_administrator/handlers"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func init() {
	log.SetFlags(log.Ldate | log.Lshortfile)
}

func dummyHandler(c *gin.Context) {
	fmt.Fprintf(c.Writer, "[%s] Response for you :)", time.Now().Format(time.RFC3339Nano))
}

func main() {

	// Logging to a file.
	f, _ := os.Create("gin.log")
	// Write the logs to file and console at the same time.
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	// Disable Console Color, you don't need console color when writing the logs to file.
	// gin.DisableConsoleColor()
	gin.ForceConsoleColor()

	// マルチプレクサにはginを使う
	// gin.New() はミドルウェアが何も入っていないまっさらなginEngineを返す
	router := gin.New()

	// Setting Logger with custom format
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())

	/* 以下ルーティング情報 */
	router.GET("/", dummyHandler)

	// 部屋一覧
	router.GET("/rooms", handlers.GetRooms)
	// 部屋を追加
	router.POST("/rooms", handlers.RegisterRoom)

	// 特定の部屋の操作
	router.GET("/rooms/:roomID", handlers.GetRoom)       // 部屋のトップページ
	router.PUT("/rooms/:roomID", handlers.UpdateRoom)    // 部屋の情報を更新
	router.DELETE("/rooms/:roomID", handlers.DeleteRoom) // 部屋を削除

	// センサー関連
	router.GET("/rooms/:roomID/sensors", handlers.GetSensors)      // 搭載されている全センサー
	router.POST("/rooms/:roomID/sensors", handlers.RegisterSensor) // センサーを追加

	// 特定のセンサーの操作
	router.GET("/rooms/:roomID/sensors/:sensorID", handlers.GetSensor)       // 特定のセンサーに関する情報
	router.PUT("/rooms/:roomID/sensors/:sensorID", handlers.UpdateSensor)    // 特定のセンサーを更新
	router.DELETE("/rooms/:roomID/sensors/:sensorID", handlers.DeleteSensor) // 特定のセンサーを削除

	// 特定のセンサーの記録
	router.GET("/rooms/:roomID/sensors/:sensorID/measurements", handlers.GetMeasurements)      // 特定のセンサーの記録を返す
	router.POST("/rooms/:roomID/sensors/:sensorID/measurements", handlers.RegisterMeasurement) // 特定のセンサーの記録を記録
	router.GET("/rooms/:roomID/sensors/:sensorID/latest", handlers.GetLatestMeasurement)       // 特定のセンサーの最新の記録を返す

	// コントローラー関係
	router.GET("/rooms/:roomID/controllers", handlers.GetAllControllersInRoom) // 利用可能な全コントローラー
	router.POST("/rooms/:roomID/controllers", handlers.RegisterController)     // コントローラーを追加

	// 特定のコントローラーの情報を取得
	router.GET("/rooms/:roomID/controllers/:controllerID", handlers.GetController) // コントローラーの情報や利用可能なコマンドを返す

	// 特定のコントローラーのコマンド
	router.GET("/rooms/:roomID/controllers/:controllerID/commands", handlers.GetControllerCommands) // コマンド一覧
	router.POST("/rooms/:roomID/controllers/:controllerID/commands", handlers.RegisterCommand)      // コマンドを追加

	// 特定のコントローラーの特定のコマンド
	router.GET("/rooms/:roomID/controllers/:controllerID/commands/:commandID", dummyHandler)    // コマンドの情報
	router.POST("/rooms/:roomID/controllers/:controllerID/commands/:commandID", dummyHandler)   // コマンド実行
	router.PUT("/rooms/:roomID/controllers/:controllerID/commands/:commandID", dummyHandler)    // コマンドを更新
	router.DELETE("/rooms/:roomID/controllers/:controllerID/commands/:commandID", dummyHandler) // コマンドを削除

	// Run server
	server := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
	log.Println("start runserver at localhost:8080")
}
