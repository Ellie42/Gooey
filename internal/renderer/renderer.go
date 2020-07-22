package renderer

import (
	"fmt"
	"github.com/go-gl/gl/v4.6-core/gl"
	"strings"
)

const (
	vertexShaderSource = `
    #version 410

	layout(location = 1) in vec2 vertexUV;

    in vec3 vp;

	out vec2 UV;

    void main() {
        gl_Position = vec4(vp, 1.0);

		UV = vertexUV;
    }
` + "\x00"

	fragmentShaderSource = `
	#version 410

	in vec2 UV;

	uniform sampler2D tex;

	out vec4 outputColor;

	void main() {
		outputColor = texture(tex, UV);
		//outputColor = vec4(UV.x, UV.y,0, 1);
		//outputColor = vec4(1,1,1,1);
	}
` + "\x00"
)

type Renderer struct {
	defaultProgramHandle uint32
}

func (r *Renderer) Init() {
	if err := gl.Init(); err != nil {
		panic(fmt.Sprintf("failed to initialise opengl"))
	}

	r.defaultProgramHandle = gl.CreateProgram()

	vertexShader, err := compileShader(vertexShaderSource, gl.VERTEX_SHADER)

	if err != nil {
		panic(err)
	}

	fragmentShader, err := compileShader(fragmentShaderSource, gl.FRAGMENT_SHADER)

	if err != nil {
		panic(err)
	}

	gl.AttachShader(r.defaultProgramHandle, vertexShader)
	gl.AttachShader(r.defaultProgramHandle, fragmentShader)
	gl.LinkProgram(r.defaultProgramHandle)

	//glMatrixMode(GL_PROJECTION);
	//glLoadIdentity();
	//glOrtho(0,width,0,height,-1,1);
	//glMatrixMode(GL_MODELVIEW);

	gl.ClearColor(0.1, 0.1, 0.2, 1.0)
}

func (r *Renderer) Clear() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	gl.UseProgram(r.defaultProgramHandle)
}

func NewRenderer() *Renderer {
	return &Renderer{}
}

func compileShader(source string, shaderType uint32) (uint32, error) {
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
