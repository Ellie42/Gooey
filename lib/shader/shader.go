package shader

import (
	"fmt"
	"github.com/go-gl/gl/v4.6-compatibility/gl"
	"strings"
)

func CompileProgram(vertSource string, fragSource string) uint32 {
	progHandle := gl.CreateProgram()

	vertHandle, err := CompileShader(vertSource, gl.VERTEX_SHADER)

	if err != nil {
		panic(err)
	}

	fragHandle, err := CompileShader(fragSource, gl.FRAGMENT_SHADER)

	if err != nil {
		panic(err)
	}

	gl.AttachShader(progHandle, vertHandle)
	gl.AttachShader(progHandle, fragHandle)
	gl.LinkProgram(progHandle)

	return progHandle
}

func CompileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	csources, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to compile %v: %v", source, log)
	}

	return shader, nil
}
