package widget

type Widget interface {
	Init(parent WidgetParent)
	Render()
}
