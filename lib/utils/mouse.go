package utils

import (
	"github.com/go-vgo/robotgo"
)

type MouseController struct{}

func NewMouseController() *MouseController {
	return &MouseController{}
}

func (this *MouseController) MoveCursor(mx, my int) {
	width, height := robotgo.GetScreenSize()
	x, y := robotgo.GetMousePos()
	x = x + mx
	if x < 0 {
		x = 0
	} else if x > width {
		x = width
	}
	y = y + my
	if y < 0 {
		y = 0
	} else if y > height {
		y = height
	}
	robotgo.MoveMouse(x, y)
}

func (this *MouseController) MouseLeftButtonDown() {
	robotgo.MouseToggle("down", "left")
}

func (this *MouseController) MouseLeftButtonUp() {
	robotgo.MouseToggle("up", "left")
}

func (this *MouseController) MouseRightButtonDown() {
	robotgo.MouseToggle("down", "right")
}

func (this *MouseController) MouseRightButtonUp() {
	robotgo.MouseToggle("up", "right")
}

func (this *MouseController) MouseScroll(cx, cy int) {
	if cy > 0 {
		robotgo.ScrollMouse(cy, "down")
	} else {
		robotgo.ScrollMouse(-cy, "up")
	}
}
