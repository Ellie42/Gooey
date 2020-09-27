package eventmanager

import (
	"git.agehadev.com/elliebelly/gooey/lib/dimension"
	"git.agehadev.com/elliebelly/gooey/pkg/draw"
	"github.com/go-gl/glfw/v3.3/glfw"
	"math"
)

type KeyEventListener func(key glfw.Key)

type EventManager struct {
	mouseClickHandlers  map[int]MouseClickHandler
	mouseScrollHandlers map[int]MouseScrollHandler
	mouseClickRectIndex int
	areaContext         EventAreaContext
	glfwWindow          *glfw.Window

	keyDownListeners     map[glfw.Key][]KeyEventListener
	keyUpListeners       map[glfw.Key][]KeyEventListener
	mouseScrollRectIndex int
}

type EventAreaContext interface {
	GetTotalArea() dimension.DimensionsFloat32
}

type RectHandler interface {
	GetZIndex() int
	GetRects() []dimension.Rect
}

type MouseClickHandler interface {
	RectHandler
	OnClick(position dimension.Vector2)
	OnMouseUp()
}

type ScrollDirection int

const (
	ScrollDirectionUp ScrollDirection = iota
	ScrollDirectionDown
)

type MouseScrollHandler interface {
	RectHandler
	OnScroll(xoff, yoff float64)
}

func (e *EventManager) HandleMouseDown() {
	mouseX, mouseY := e.glfwWindow.GetCursorPos()

	foundIndex, foundCollision := e.findHandledWidgetAtPosition(func(i int) RectHandler {
		return e.mouseClickHandlers[i]
	}, len(e.mouseClickHandlers), mouseX, mouseY)

	if !foundCollision {
		return
	}

	e.mouseClickHandlers[foundIndex].OnClick(dimension.Vector2{
		float32(mouseX), float32(mouseY),
	})
}

func (e *EventManager) HandleMouseScroll(xoff float64, yoff float64) {
	mouseX, mouseY := e.glfwWindow.GetCursorPos()

	foundIndex, foundCollision := e.findHandledWidgetAtPosition(func(i int) RectHandler {
		return e.mouseScrollHandlers[i]
	}, len(e.mouseScrollHandlers), mouseX, mouseY)

	if !foundCollision {
		return
	}

	e.mouseScrollHandlers[foundIndex].OnScroll(xoff, yoff)
}

func (e *EventManager) findHandledWidgetAtPosition(handlerProvider func(i int) RectHandler, handlerCount int, mouseX float64, mouseY float64) (int, bool) {
	lowestZIndex := math.MaxInt64
	foundIndex := -1
	foundCollision := false

	res := e.areaContext.GetTotalArea()

	mouseY = float64(res.Height) - mouseY

	screenRect := dimension.Rect{
		res.Width, res.Height, res.Width, res.Height,
	}

	for handlerIndex := 0; handlerIndex < handlerCount; handlerIndex++ {
		rectHandler := handlerProvider(handlerIndex)
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

	return foundIndex, foundCollision
}

func (e *EventManager) RegisterMouseScrollHandler(d MouseScrollHandler) (index int) {
	index = e.mouseScrollRectIndex

	e.mouseScrollHandlers[e.mouseScrollRectIndex] = d

	e.mouseScrollRectIndex++

	return
}

func (e *EventManager) RegisterMouseClickHandler(d MouseClickHandler) (index int) {
	index = e.mouseClickRectIndex

	e.mouseClickHandlers[e.mouseClickRectIndex] = d

	e.mouseClickRectIndex++

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
	for _, handler := range e.mouseClickHandlers {
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
		mouseClickHandlers:  make(map[int]MouseClickHandler, 0),
		mouseScrollHandlers: make(map[int]MouseScrollHandler, 0),
		keyDownListeners:    make(map[glfw.Key][]KeyEventListener),
		keyUpListeners:      make(map[glfw.Key][]KeyEventListener),
	}
}
