package draw

import (
	"git.agehadev.com/elliebelly/gooey/lib/dimension"
	"github.com/go-gl/gl/v4.6-core/gl"
)

var lineVAO uint32
var lineVBO uint32
var lineColourVBO uint32

var pointsBuffer []dimension.Vector3
var coloursBuffer []RGBA

func Line(linePoints []dimension.Vector3, colours []RGBA) {
	if lineVAO == 0 {
		pointsBuffer = make([]dimension.Vector3, 256)
		coloursBuffer = make([]RGBA, 256)

		lineVAO = genVAO(1)[0]
		gl.BindVertexArray(lineVAO)

		vbos := genVBO(2)

		lineVBO, lineColourVBO = vbos[0], vbos[1]

		configureVBOs()
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
			3* // Number of Float32s per Vector3
			len(pointsBuffer),
		gl.Ptr(pointsBuffer), gl.STATIC_DRAW)

	gl.DrawArrays(gl.LINES, 0, int32(len(linePoints)))
}

func configureVBOs() {
	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, lineVBO)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)

	gl.EnableVertexAttribArray(1)
	gl.BindBuffer(gl.ARRAY_BUFFER, lineColourVBO)
	gl.VertexAttribPointer(1, 4, gl.FLOAT, false, 0, nil)
}
