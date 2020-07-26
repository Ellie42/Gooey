package dimension

type BoundingBox struct {
	MinX, MinY, MaxX, MaxY float32
}

func (b BoundingBox) Contains(x float32, y float32) bool {
	return x > b.MinX && x < b.MaxX && y > b.MinY && y < b.MaxY
}

func (b BoundingBox) ToRect() Rect {
	return Rect{
		X:      b.MinX,
		Y:      b.MinY,
		Width:  b.MaxX - b.MinX,
		Height: b.MaxY - b.MinY,
	}
}

type DirectionalRect struct {
	Top, Right, Bottom, Left float32
}

type Rect struct {
	X, Y, Width, Height float32
}

func (r Rect) RelativeTo(parent Rect) Rect {
	r.X = parent.X + r.X*float32(parent.Width)
	r.Y = parent.Y + r.Y*float32(parent.Height)
	r.Width = r.Width * float32(parent.Width)
	r.Height = r.Height * float32(parent.Height)
	return r
}

func (r Rect) WithPadding(padding DirectionalRect) Rect {
	r.X += padding.Left * r.Width
	r.Y += padding.Bottom * r.Height
	r.Width -= padding.Right + padding.Left*r.Width
	r.Height -= padding.Top + padding.Bottom*r.Height
	return r
}

func (r Rect) MultipliedBy(rect Rect) Rect {
	r.X *= rect.X
	r.Y *= rect.Y
	r.Width *= rect.Width
	r.Height *= rect.Height
	return r
}

func (r Rect) ToBoundingBox() BoundingBox {
	return BoundingBox{
		MinX: r.X,
		MinY: r.Y,
		MaxX: r.X + r.Width,
		MaxY: r.Y + r.Height,
	}
}

func (r Rect) GetRatioX() float32 {
	return r.Height / r.Width
}

func (r Rect) GetRatioY() float32 {
	return r.Width / r.Height
}

func (r Rect) MultipliedByDimension(resolution DimensionsInt) Rect {
	r.X *= float32(resolution.Width)
	r.Y *= float32(resolution.Height)
	r.Width *= float32(resolution.Width)
	r.Height *= float32(resolution.Height)
	return r
}

func (r Rect) Scale(amount float32) Rect {
	r.Width *= amount
	r.Height *= amount
	return r
}

func (r Rect) Shrink(amount float32) Rect {
	r.X += amount/2
	r.Y += amount/2
	r.Width -= amount
	r.Height -= amount
	return r
}
