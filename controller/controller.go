package controller

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/tosashimanto/iot-camera-web/controller/handler"
)

var Server *echo.Echo

// init echo web server
func Init() {

	Server = echo.New()
	// Middlewareセット
	Server.Use(middleware.Logger())
	Server.Use(middleware.Recover())

	// 静的ファイルパス設定
	// Front
	Server.Static("/*", "./static/front")

	// http://localhost:8080/d3_map/us.json
	Server.Static("/d3_map/*", "./static/map")
	Server.Static("/d3_graph/*", "./static/data")
	// http://localhost:8080/cors_test/websockete_test_chat_home.html
	Server.Static("/cors_test/*", "./static/cors_test")
	// http://localhost:8080/tmp/ID00000001_20190119_214410.jpgでアクセス
	Server.Static("/tmp/*", "./tmp/")

	// CORSセット
	Server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	api := Server.Group("/iot_camera/v1")

	// Tokenチェック
	api.POST("/auth/checkToken", handler.CheckFirebaseToken)

	// URL取得
	api.GET("/getUploadURL", handler.GetUplaodURL)

	// 擬似S3 uplaod
	api.PUT("/upload/:index", handler.UploadImage)

	// 画像List
	api.GET("/getImageList", handler.GetImageList)
}
