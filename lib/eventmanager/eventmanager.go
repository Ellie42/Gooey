package eventmanager

import (
	"fmt"
	"git.agehadev.com/elliebelly/gooey/lib/dimension"
	"git.agehadev.com/elliebelly/gooey/pkg/draw"
	"github.com/go-gl/glfw/v3.3/glfw"
	"math"
)

type KeyEventListener func(key glfw.Key)

type EventManager struct {
	mouseRectHandlers map[int]MouseRectClickHandler
	rectIndex         int
	areaContext       EventAreaContext
	glfwWindow        *glfw.Window

	keyDownListeners map[glfw.Key][]KeyEventListener
	keyUpListeners   map[glfw.Key][]KeyEventListener
}

type EventAreaContext interface {
	GetTotalArea() dimension.DimensionsFloat32
}

type MouseRectClickHandler interface {
	GetZIndex() int
	GetRects() []dimension.Rect
	OnClick(position dimension.Vector2)
	OnMouseUp()
}

func (e *EventManager) HandleMouseDown() {
	lowestZIndex := math.MaxInt64
	foundIndex := 0
	foundCollision := false

	res := e.areaContext.GetTotalArea()

	mouseX, mouseY := e.glfwWindow.GetCursorPos()

	mouseY = float64(res.Height) - mouseY

	screenRect := dimension.Rect{
		res.Width, res.Height, res.Width, res.Height,
	}

	fmt.Printf("%f %f - click\n", mouseX, mouseY)

	for handlerIndex, rectHandler := range e.mouseRectHandlers {
		zIndex := rectHandler.GetZIndex()
		rects := rectHandler.GetRects()

		if zIndex > lowestZIndex {
			continue
		}

		for _, rect := range rects {
			boundingBox := rect.MultipliedBy(screenRect).ToBoundingBox()

			draw.SquareEdge(boundingBox.ToRect().Shrink(0.1), draw.Red)

			if boundingBox.Contains(float32(mouseX), float32(mouseY)) {
				lowestZIndex = zIndex
				foundCollision = true
				foundIndex = handlerIndex
				break
			}
		}
	}

	if !foundCollision {
		return
	}

	e.mouseRectHandlers[foundIndex].OnClick(dimension.Vector2{
		float32(mouseX), float32(mouseY),
	})
}

func (e *EventManager) RegisterRect(d MouseRectClickHandler) (index int) {
	index = e.rectIndex

	e.mouseRectHandlers[e.rectIndex] = d

	e.rectIndex++

	return
}

func (e *EventManager) Init(areaContext EventAreaContext, window2 *glfw.Window) {
	e.areaContext = areaContext
	e.glfwWindow = window2
}

func (e *EventManager) MousePosition() dimension.Vector2 {
	posX, posY := e.glfwWindow.GetCursorPos()

	return dimension.Vector2{
		float32(posX),
		e.areaContext.GetTotalArea().Height - float32(posY),
	}
}

func (e *EventManager) HandleMouseClickCollisions() {
	for _, handler := range e.mouseRectHandlers {
		handler.OnMouseUp()
	}
}

func (e *EventManager) PixelPositionToRelative(position dimension.Vector2) dimension.Vector2 {
	pixelArea := e.areaContext.GetTotalArea()
	position.X /= pixelArea.Width
	position.Y /= pixelArea.Height
	return position
}

func (e *EventManager) HandleKeyDown(key glfw.Key) {
	if listeners, ok := e.keyDownListeners[key]; ok {
		for _, listener := range listeners {
			listener(key)
		}
	}
}

func (e *EventManager) HandleKeyUp(key glfw.Key) {
	if listeners, ok := e.keyUpListeners[key]; ok {
		for _, listener := range listeners {
			listener(key)
		}
	}
}

func (e *EventManager) AddOnKeyDownListener(key glfw.Key, listener KeyEventListener) {
	listeners, ok := e.keyDownListeners[key]

	if !ok {
		e.keyDownListeners[key] = make([]KeyEventListener, 0)
	}

	e.keyDownListeners[key] = append(listeners, listener)
}

func (e *EventManager) AddOnKeyUpListener(key glfw.Key, listener KeyEventListener) {
	listeners, ok := e.keyUpListeners[key]

	if !ok {
		e.keyUpListeners[key] = make([]KeyEventListener, 0)
	}

	e.keyUpListeners[key] = append(listeners, listener)
}

func NewEventManager() *EventManager {
	return &EventManager{
		mouseRectHandlers: make(map[int]MouseRectClickHandler, 0),
		keyDownListeners:  make(map[glfw.Key][]KeyEventListener),
		keyUpListeners:    make(map[glfw.Key][]KeyEventListener),
	}
}
