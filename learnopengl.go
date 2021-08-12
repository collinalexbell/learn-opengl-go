package main // import "github.com/go-gl/example/gl41core-cube"

import (
	"runtime"
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func clearToCyberpunkColor() {
	gl.ClearColor(4.0/255.0, 217.0/255.0, 255.0/255.0, 255.0/255.0)
	gl.Clear(gl.COLOR_BUFFER_BIT)
}

func processInput(window *glfw.Window) {
	if(window.GetKey(glfw.KeyEscape) == glfw.Press) {
		window.SetShouldClose(true)
	}
}

func init() {
	runtime.LockOSThread()
}

func main() {
	err := glfw.Init()
	if err != nil {
		panic(err)
	}
	defer glfw.Terminate()

	err = gl.Init()

	if err != nil {
		panic(err)
	}

	window, err := glfw.CreateWindow(800, 600, "LearnOpenGL", nil, nil)
	if err != nil {
		panic(err)
	}

	window.MakeContextCurrent()
	gl.Viewport(0, 0, 800, 600)

	window.SetFramebufferSizeCallback(func(w *glfw.Window, width int, height int) {
		gl.Viewport(0, 0, int32(width), int32(height))
	})

	vertices := []float32{
		-0.5, -0.5, 0.0,
		0.5, -0.5, 0.0,
		0.0, 0.5, 0.0,
	}

	var VAO uint32
	gl.GenVertexArrays(1, &VAO)
	gl.BindVertexArray(VAO)

	var VBO uint32
	gl.GenBuffers(1, &VBO)
	gl.BindBuffer(gl.ARRAY_BUFFER, VBO)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices) * 4, gl.Ptr(vertices), gl.STATIC_DRAW)

	shaderSource := gl.Str("#version 330 core\n" +
	"layout (location = 0) in vec3 aPos;\n" +
	"void main()\n" +
	"{\n" +
	" gl_Position = vec4(aPos.x, aPos.y, aPos.z, 1.0);\n" +
	"}\x00")

	vertexShader := gl.CreateShader(gl.VERTEX_SHADER)
	gl.ShaderSource(vertexShader, 1, &shaderSource, nil)
	gl.CompileShader(vertexShader)

	var success int32
	gl.GetShaderiv(vertexShader, gl.COMPILE_STATUS, &success)

	if(success == 0) {
		panic("shader couldn't compile")
	}

	fragShaderSource := gl.Str(`
	#version 330 core
	out vec4 FragColor;
	void main()
	{
		FragColor = vec4(1.0f, 0.5f, 0.2f, 1.0f);
	}` + "\x00")

	fragShader := gl.CreateShader(gl.FRAGMENT_SHADER)
	gl.ShaderSource(fragShader, 1, &fragShaderSource, nil)
	gl.CompileShader(fragShader)

	shaderProg := gl.CreateProgram()
	gl.AttachShader(shaderProg, vertexShader)
	gl.AttachShader(shaderProg, fragShader)
	gl.LinkProgram(shaderProg)
	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragShader)

	gl.VertexAttribPointer(0,3, gl.FLOAT, false, 12, gl.PtrOffset(0))
	gl.EnableVertexAttribArray(0)

	for !window.ShouldClose() {
		// Do OpenGL stuff.
		processInput(window)
		clearToCyberpunkColor()
		gl.UseProgram(shaderProg)
		gl.BindVertexArray(VAO)
		gl.DrawArrays(gl.TRIANGLES, 0, 3)
		window.SwapBuffers()
		glfw.PollEvents()
	}
}
