package main

import (
	"mouseCtrl/lib/controller"
	"mouseCtrl/lib/handler"
	"mouseCtrl/lib/web"
)

const WebPort int = 8888
const WsPath string = "/mouse"
const WsPort int = 7777

func main() {
	go func() {
		web.StartAgent(WebPort)
	}()

	mouseController := controller.NewMouseController()
	mouseEventHandler := handler.NewMouseEventHandler(mouseController, WsPath, WsPort)
	mouseEventHandler.Start()
}
