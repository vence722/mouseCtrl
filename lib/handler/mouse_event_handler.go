package handler

import (
	"fmt"
	"mouseCtrl/lib/controller"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
	"github.com/vence722/convert"
)

type MouseEventHandler struct {
	mouseController *controller.MouseController
	path            string
	port            int
}

func NewMouseEventHandler(mouseController *controller.MouseController, path string, port int) *MouseEventHandler {
	return &MouseEventHandler{
		mouseController: mouseController,
		path:            path,
		port:            port,
	}
}

func (this *MouseEventHandler) Start() {
	http.HandleFunc(this.path, this.handler)
	http.ListenAndServe(":"+convert.Int2Str(this.port), nil)
}

func (this *MouseEventHandler) handler(rsp http.ResponseWriter, req *http.Request) {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	ws, err := upgrader.Upgrade(rsp, req, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	ch := make(chan string, 150)

	// create handle function
	go func() {
		for {
			// receive data
			recv := <-ch

			fmt.Println("recv=", recv)

			// handle data
			if "ld" == recv {
				this.mouseController.MouseLeftButtonDown()
			} else if "lu" == recv {
				this.mouseController.MouseLeftButtonUp()
			} else if "rd" == recv {
				this.mouseController.MouseRightButtonDown()
			} else if "ru" == recv {
				this.mouseController.MouseRightButtonUp()
			} else if "exit" == recv {
				break
			} else if strings.HasPrefix(recv, "sc,") {
				cx, cy := getScrollAmount(recv)
				this.mouseController.MouseScroll(cx, cy)
			} else {
				mx, my := getPoint(recv)
				this.mouseController.MoveCursor(mx, my)
			}
		}
		close(ch)
	}()

	for {
		_, data, err := ws.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}
		recv := string(data)
		ch <- recv
	}
}

func getPoint(point string) (int, int) {
	i := strings.Index(point, ",")
	if i == -1 {
		return 0, 0
	}
	x := convert.Str2Float32(point[:i])
	y := convert.Str2Float32(point[i+1:])
	return int(x), int(y)
}

func getScrollAmount(command string) (int, int) {
	arr := strings.Split(command, ",")
	if len(arr) == 0 {
		return 0, 0
	}
	return int(convert.Str2Float32(arr[1])), int(convert.Str2Float32(arr[2]))
}
