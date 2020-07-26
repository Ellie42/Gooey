package draw

type RGBA struct {
	R, G, B, A float32
}

var (
	Red   RGBA = RGBA{1, 0, 0, 1}
	Green RGBA = RGBA{0, 1, 1, 1}
	Blue  RGBA = RGBA{0, 0, 1, 1}
	White RGBA = RGBA{1, 1, 1, 1}
)
