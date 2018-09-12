package main

import (
	"mouseCtrl/lib/handler"
	"mouseCtrl/lib/utils"
	"mouseCtrl/lib/web"
)

const WebPort int = 8888
const WsPath string = "/mouse"
const WsPort int = 7777

func main() {
	go func() {
		web.StartAgent(WebPort)
	}()

	mouseController := utils.NewMouseController()
	mouseEventHandler := handler.NewMouseEventHandler(mouseController, WsPath, WsPort)
	mouseEventHandler.Start()
}
