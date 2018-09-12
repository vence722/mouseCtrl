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
	http.HandleFunc("/", mouseCtrlHandler)
	http.HandleFunc("/auth", authHandler)
	http.HandleFunc("/token", tokenHandler)
	http.ListenAndServe(":"+convert.Int2Str(port), nil)
}

func mouseCtrlHandler(rsp http.ResponseWriter, req *http.Request) {
	f, _ := os.Open("lib/web/mouse_ctrl.html")
	data, _ := ioutil.ReadAll(f)
	rsp.Write(data)
}

func authHandler(rsp http.ResponseWriter, req *http.Request) {
	f, _ := os.Open("lib/web/auth.html")
	data, _ := ioutil.ReadAll(f)
	rsp.Write(data)
}

func tokenHandler(rsp http.ResponseWriter, req *http.Request) {
	_, qrCode, err := utils.GenerateAuthTokenAndQRCodePair()
	if err != nil {
		rsp.Write([]byte(err.Error()))
		return
	}
	rsp.Write([]byte(base64.StdEncoding.EncodeToString(qrCode)))
}
