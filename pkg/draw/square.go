package draw

import (
	"git.agehadev.com/elliebelly/gooey/lib/dimension"
	"github.com/go-gl/gl/v4.6-core/gl"
)

var quadTriangles = []float32{
	0, 1, 0, // Top Left
	0, 0, 0, // Bottom Left
	1, 0, 0, // Bottom Right

	0, 1, 0, // Top Left
	1, 0, 0, // Bottom Right
	1, 1, 0, // Top Right
}

var quadUVs = []float32{
	0, 0, // Top Left
	0, 1, // Bottom Left
	1, 1, // Bottom Right

	0, 0, // Top Left
	1, 1, // Bottom Right
	1, 0, // Top Right
}

var squareVertVBO uint32
var squareUVVBO uint32
var squareVertColourVBO uint32
var squareVAO uint32
var squareTexture uint32
var squareVertBuffer = make([]dimension.Vector3, 6)
var squareUVBuffer = make([]dimension.Vector2, 6)

func SquareFilled(rect dimension.Rect, colour RGBA) {
	colours := make([]RGBA, 6)

	for i := 0; i < 6; i++ {
		colours[i] = colour
	}

	if squareVAO == 0 {
		squareVAO = genVAO(1)[0]
		gl.BindVertexArray(squareVAO)

		vbos := genVBO(3)
		squareVertVBO, squareVertColourVBO, squareUVVBO = vbos[0], vbos[1], vbos[2]

		gl.EnableVertexAttribArray(0)
		gl.BindBuffer(gl.ARRAY_BUFFER, squareVertVBO)
		gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)

		gl.EnableVertexAttribArray(1)
		gl.BindBuffer(gl.ARRAY_BUFFER, squareVertColourVBO)
		gl.VertexAttribPointer(1, 4, gl.FLOAT, false, 0, nil)

		gl.EnableVertexAttribArray(2)
		gl.BindBuffer(gl.ARRAY_BUFFER, squareUVVBO)
		gl.VertexAttribPointer(2, 2, gl.FLOAT, false, 0, nil)

		gl.GenTextures(1, &squareTexture)
		gl.ActiveTexture(gl.TEXTURE0)
	}

	copy(squareVertBuffer, preparePositionsForGL([]dimension.Vector3{
		{rect.X, rect.Y + rect.Height, 1},
		{rect.X, rect.Y, 1},
		{rect.X + rect.Width, rect.Y, 1},

		{rect.X, rect.Y + rect.Height, 1},
		{rect.X + rect.Width, rect.Y, 1},
		{rect.X + rect.Width, rect.Y + rect.Height, 1},
	}))

	gl.BindVertexArray(squareVAO)

	gl.BindBuffer(gl.ARRAY_BUFFER, squareVertVBO)
	gl.BufferData(gl.ARRAY_BUFFER, 4*3*len(squareVertBuffer), gl.Ptr(squareVertBuffer), gl.STATIC_DRAW)

	gl.BindBuffer(gl.ARRAY_BUFFER, squareVertColourVBO)
	gl.BufferData(gl.ARRAY_BUFFER, 4*4*len(colours), gl.Ptr(colours), gl.STATIC_DRAW)

	gl.BindBuffer(gl.ARRAY_BUFFER, squareUVVBO)
	gl.BufferData(gl.ARRAY_BUFFER, 4*2*len(squareUVBuffer), gl.Ptr(squareUVBuffer), gl.STATIC_DRAW)

	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(quadTriangles)/3))
}

func SquareEdge(rect dimension.Rect, colour RGBA) {
	colours := make([]RGBA, 8)

	for i := 0; i < 8; i++ {
		colours[i] = colour
	}

	Line(preparePositionsForGL([]dimension.Vector3{
		{rect.X, rect.Y, 1},
		{rect.X, rect.Y + rect.Height, 1},
		{rect.X, rect.Y + rect.Height, 1},
		{rect.X + rect.Width, rect.Y + rect.Height, 1},
		{rect.X + rect.Width, rect.Y + rect.Height, 1},
		{rect.X + rect.Width, rect.Y, 1},
		{rect.X + rect.Width, rect.Y, 1},
		{rect.X, rect.Y, 1},
	}), colours)
}

func preparePositionsForGL(vectors []dimension.Vector3) []dimension.Vector3 {
	for i, _ := range vectors {
		vectors[i].X = vectors[i].X*2 - 1
		vectors[i].Y = vectors[i].Y*2 - 1
		//vectors[i].X = vectors[i].X
		//vectors[i].Y = vectors[i].Y
	}

	return vectors
}
