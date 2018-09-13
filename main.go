package main

import (
	"mouseCtrl/lib/web"
)

const WebPort int = 8888
const WsPath string = "/mouse"
const WsPort int = 7777

func main() {
	web.StartAgent(WebPort, WsPath, WsPort)
}
