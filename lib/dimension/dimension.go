package dimension

type DimensionsInt struct {
	Width, Height int
}

type DimensionsFloat32 struct {
	Width, Height float32
}

type Dimensions struct {
	Width, Height *Size
}

type Dimension int

const (
	Width Dimension = iota
	Height
)

func (bounds Dimensions) Bind(rect *Rect, parentRect Rect, dimension Dimension, resolution DimensionsInt) {
	var pixelDimensionValue, parentRectDimensionValue float32
	var windowResolutionDimensionValue int
	var size *Size
	var value *float32

	if rect == nil {
		panic("bind rect dimension failed, rect cannot be nil")
	}

	pixelPos := rect.RelativeToAbsolute(parentRect).MultipliedByDimension(resolution)

	switch dimension {
	case Width:
		pixelDimensionValue = pixelPos.Width
		parentRectDimensionValue = parentRect.Width
		windowResolutionDimensionValue = resolution.Width
		size = bounds.Width
		value = &rect.Width
	case Height:
		pixelDimensionValue = pixelPos.Height
		parentRectDimensionValue = parentRect.Height
		windowResolutionDimensionValue = resolution.Height
		size = bounds.Height
		value = &rect.Height
	}

	if size == nil || value == nil {
		return
	}

	switch size.Unit {
	case SizeUnitPixels:
		if pixelDimensionValue > size.Amount {
			*value = size.Amount / float32(windowResolutionDimensionValue) / parentRectDimensionValue
		}
	case SizeUnitRatio:
		if *value > size.Amount {
			*value = size.Amount
		}
	}
}
