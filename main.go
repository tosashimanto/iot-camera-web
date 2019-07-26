package main

import (
	"fmt"
	"github.com/tosashimanto/iot-camera-web/controller"
	"github.com/tosashimanto/iot-camera-web/db/gorm"
	"os"
	"strings"
)

func main() {

	if len(os.Args) > 1 {
		fmt.Println("", os.Args[1])
	} else {
		fmt.Println("Args len=0")
	}
	KEY := os.Getenv("PORT")
	fmt.Println("KEY=", KEY)
	fmt.Println()
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Println(pair[0])
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		fmt.Println("$PORT must be set")
	}

	// init server
	controller.Init()

	gorm.Init()

	// run server
	fmt.Println("run server")
	controller.Server.Logger.Fatal(controller.Server.Start(":" + port))

	// Websocket
	//fmt.Println("Websocket")
	//websocket.StartWebSocket()

}
