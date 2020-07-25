package draw

import "git.agehadev.com/elliebelly/gooey/lib/dimension"

func Square(rect dimension.Rect) {
	Line(preparePositionsForGL([]dimension.Vector2{
		{rect.X, rect.Y},
		{rect.X, rect.Y + rect.Height},
		{rect.X, rect.Y + rect.Height},
		{rect.X + rect.Width, rect.Y + rect.Height},
		{rect.X + rect.Width, rect.Y + rect.Height},
		{rect.X + rect.Width, rect.Y},
		{rect.X + rect.Width, rect.Y},
		{rect.X, rect.Y},
	}))
}

func preparePositionsForGL(vector2s []dimension.Vector2) []dimension.Vector2 {
	for i, _ := range vector2s {
		vector2s[i].X = vector2s[i].X*2 - 1
		vector2s[i].Y = vector2s[i].Y*2 - 1
	}

	return vector2s
}
