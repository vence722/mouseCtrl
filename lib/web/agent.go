package web

import (
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/vence722/convert"

	"mouseCtrl/lib/utils"
)

// authentication token
var AuthToken string = ""

// connected flag
var IsConnected bool = false

func StartAgent(port int) {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/qrCode", qrCodeHandler(port))
	http.HandleFunc("/auth", authHandler)
	http.HandleFunc("/control", mouseCtrlHandler)
	http.ListenAndServe(":"+convert.Int2Str(port), nil)
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
