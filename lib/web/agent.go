package web

import (
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/vence722/convert"

	"mouseCtrl/lib/handler"
	"mouseCtrl/lib/utils"
)

// mouse event controller
var mouseEventHandler *handler.MouseEventHandler

// authentication token
var AuthToken string = ""

// connected flag
var IsConnected bool = false

func StartAgent(webPort int, wsPath string, wsPort int) {
	// start mouse controller
	go func() {
		mouseController := utils.NewMouseController()
		mouseEventHandler = handler.NewMouseEventHandler(mouseController, wsPath, wsPort)
		mouseEventHandler.Start()
	}()

	// start web agent
	go func() {
		http.HandleFunc("/", indexHandler)
		http.HandleFunc("/qrCode", qrCodeHandler(webPort))
		http.HandleFunc("/auth", authHandler)
		http.HandleFunc("/disconnect", disconnectHandler)
		http.HandleFunc("/control", mouseCtrlHandler)
		http.ListenAndServe(":"+convert.Int2Str(webPort), nil)
	}()

	forever := make(chan bool)
	<-forever
}

func indexHandler(rsp http.ResponseWriter, req *http.Request) {
	f, _ := os.Open("lib/web/index.html")
	data, _ := ioutil.ReadAll(f)
	rsp.Write(data)
}

func qrCodeHandler(port int) func(rsp http.ResponseWriter, req *http.Request) {
	return func(rsp http.ResponseWriter, req *http.Request) {
		qrCode, authToken, err := utils.GenerateQRCodeAndAuthToken(port)
		if err != nil {
			rsp.Write([]byte(err.Error()))
			return
		}
		// set auth token
		AuthToken = authToken
		rsp.Write([]byte(base64.StdEncoding.EncodeToString(qrCode)))
	}
}

func authHandler(rsp http.ResponseWriter, req *http.Request) {
	if IsConnected {
		rsp.Write([]byte("ok"))
	} else {
		rsp.Write([]byte("no"))
	}
}

func disconnectHandler(rsp http.ResponseWriter, req *http.Request) {
	if IsConnected {
		AuthToken = ""
		IsConnected = false
		mouseEventHandler.CloseCurrentConn()
	}
	rsp.Write([]byte("ok"))
}

func mouseCtrlHandler(rsp http.ResponseWriter, req *http.Request) {
	// verify token
	token := req.URL.Query().Get("token")
	if token != AuthToken {
		rsp.Write([]byte("Authentication failed"))
		rsp.WriteHeader(400)
		return
	}

	IsConnected = true
	f, _ := os.Open("lib/web/mouse_ctrl.html")
	data, _ := ioutil.ReadAll(f)
	rsp.Write(data)
}
