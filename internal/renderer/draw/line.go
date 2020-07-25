package draw

import (
	"git.agehadev.com/elliebelly/gooey/lib/dimension"
	"github.com/go-gl/gl/v4.6-core/gl"
	"unsafe"
)

var lineVBO uint32
var lineVAO uint32

var pointsBuffer []dimension.Vector2
var pointsBufferPointer unsafe.Pointer

func Line(linePoints []dimension.Vector2) {
	if lineVBO == 0 {
		pointsBuffer = make([]dimension.Vector2, 256)
		genLineVBO(pointsBuffer)
	}

	if lineVAO == 0 {
		genLineVAO()
	}

	copy(pointsBuffer[0:], linePoints)

	gl.BindVertexArray(lineVAO)
	gl.BindBuffer(gl.ARRAY_BUFFER, lineVBO)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(pointsBuffer)*2, gl.Ptr(pointsBuffer), gl.STATIC_DRAW)
	gl.DrawArrays(gl.LINES, 0, int32(len(linePoints)*2))
}

func genLineVAO() {
	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)
	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointer(0, 2, gl.FLOAT, false, 0, nil)
	lineVAO = vao
}

func genLineVBO(points []dimension.Vector2) {
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	lineVBO = vbo
}
