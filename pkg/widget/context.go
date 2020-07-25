package widget

import (
	"git.agehadev.com/elliebelly/gooey/lib/dimension"
	"git.agehadev.com/elliebelly/gooey/lib/eventmanager"
	"git.agehadev.com/elliebelly/gooey/lib/input"
	"github.com/go-gl/glfw/v3.3/glfw"
)

type WindowContext struct {
	Resolution   dimension.DimensionsInt
	Rect         dimension.Rect
	Input        input.Input
	EventManager eventmanager.EventManager
	GLFWWindow   *glfw.Window
	ShowDebug    bool
}

var Context *WindowContext
