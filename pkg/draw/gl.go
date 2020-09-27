package draw

import "github.com/go-gl/gl/v4.6-compatibility/gl"

func genVAO(n int32) []uint32 {
	vaos := make([]uint32, n)

	gl.GenVertexArrays(n, &vaos[0])

	return vaos
}

func genVBO(n int32) []uint32 {
	vbos := make([]uint32, n)

	gl.GenBuffers(n, &vbos[0])

	return vbos
}
