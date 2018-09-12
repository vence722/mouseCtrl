package web

import (
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/vence722/convert"

	"mouseCtrl/lib/utils"
)

func StartAgent(port int) {
	http.HandleFunc("/", authHandler)
	http.HandleFunc("/qrCode", qrCodeHandler(port))
	http.HandleFunc("/control", mouseCtrlHandler)
	http.ListenAndServe(":"+convert.Int2Str(port), nil)
}

func authHandler(rsp http.ResponseWriter, req *http.Request) {
	f, _ := os.Open("lib/web/auth.html")
	data, _ := ioutil.ReadAll(f)
	rsp.Write(data)
}

func qrCodeHandler(port int) func(rsp http.ResponseWriter, req *http.Request) {
	return func(rsp http.ResponseWriter, req *http.Request) {
		qrCode, err := utils.GenerateQRCode(port)
		if err != nil {
			rsp.Write([]byte(err.Error()))
			return
		}
		rsp.Write([]byte(base64.StdEncoding.EncodeToString(qrCode)))
	}
}

func mouseCtrlHandler(rsp http.ResponseWriter, req *http.Request) {
	f, _ := os.Open("lib/web/mouse_ctrl.html")
	data, _ := ioutil.ReadAll(f)
	rsp.Write(data)
}
