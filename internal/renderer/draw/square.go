package draw

import (
	"git.agehadev.com/elliebelly/gooey/lib/dimension"
)

func SquareEdge(rect dimension.Rect, colour RGBA) {
	colours := make([]RGBA, 8)

	for i := 0; i < 8; i++ {
		colours[i] = colour
	}

	Line(preparePositionsForGL([]dimension.Vector2{
		{rect.X, rect.Y},
		{rect.X, rect.Y + rect.Height},
		{rect.X, rect.Y + rect.Height},
		{rect.X + rect.Width, rect.Y + rect.Height},
		{rect.X + rect.Width, rect.Y + rect.Height},
		{rect.X + rect.Width, rect.Y},
		{rect.X + rect.Width, rect.Y},
		{rect.X, rect.Y},
	}), colours)
}

func preparePositionsForGL(vector2s []dimension.Vector2) []dimension.Vector2 {
	for i, _ := range vector2s {
		vector2s[i].X = vector2s[i].X*2 - 1
		vector2s[i].Y = vector2s[i].Y*2 - 1
	}

	return vector2s
}
