package draw

import (
	"git.agehadev.com/elliebelly/gooey/lib/dimension"
	"github.com/go-gl/gl/v4.6-core/gl"
)

var lineVAO uint32
var lineVBO uint32
var lineColourVBO uint32

var pointsBuffer []dimension.Vector2
var coloursBuffer []RGBA

func Line(linePoints []dimension.Vector2, colours []RGBA) {
	if lineVBO == 0 {
		pointsBuffer = make([]dimension.Vector2, 256)
		coloursBuffer = make([]RGBA, 256)
		genLineVBO()
	}

	if lineVAO == 0 {
		genLineVAO()
	}

	copy(pointsBuffer[0:], linePoints)
	copy(coloursBuffer[0:], colours)

	gl.BindVertexArray(lineVAO)

	gl.BindBuffer(gl.ARRAY_BUFFER, lineColourVBO)
	gl.BufferData(gl.ARRAY_BUFFER,
		4* // float32 bytes
			4* // Number of uint8s per color.RGBA
			len(colours),
		gl.Ptr(coloursBuffer), gl.STATIC_DRAW)

	gl.BindBuffer(gl.ARRAY_BUFFER, lineVBO)
	gl.BufferData(gl.ARRAY_BUFFER,
		4* // Float32 bytes
			2* // Number of Float32s per Vector2
			len(pointsBuffer),
		gl.Ptr(pointsBuffer), gl.STATIC_DRAW)

	gl.DrawArrays(gl.LINES, 0, int32(len(linePoints)*2))
}

func genLineVAO() {
	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)

	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, lineVBO)
	gl.VertexAttribPointer(0, 2, gl.FLOAT, false, 0, nil)

	gl.EnableVertexAttribArray(1)
	gl.BindBuffer(gl.ARRAY_BUFFER, lineColourVBO)
	gl.VertexAttribPointer(1, 4, gl.FLOAT, false, 0, nil)

	lineVAO = vao
}

func genLineVBO() {
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	lineVBO = vbo

	var vbo2 uint32
	gl.GenBuffers(1, &vbo2)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo2)
	lineColourVBO = vbo2
}
