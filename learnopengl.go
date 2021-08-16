package main

import (
	"fmt"
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

// This is the go runtime init that runs before main
func init() {
	runtime.LockOSThread()
}

func newWindow() *glfw.Window {
	fmt.Println("init Window")
	err := glfw.Init()
	if err != nil {
		panic(err)
	}

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

	return window
}

func bindVerts(vertices []float32) uint32 {
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER,
		len(vertices) * 4,
		gl.Ptr(vertices),
		gl.STATIC_DRAW)
	return vbo
}

func triangleShaderProg() uint32 {
	// location = 1 selects the 2nd attribute pointer (index 1) in the VAO
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

	return shaderProg

}

func configureVAO(EBO uint32) {
	// Vertex attribute configration, stored in the VAO
	// The first attribute pointer (reference location = {0, 1} in shader)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 12, gl.PtrOffset(0))
	// Enable the configuration, stored in the VAO
	// required to use the attribpointers in the shader progs
	gl.EnableVertexAttribArray(0)
}

func createEBO(indices []uint32) uint32 {
	var EBO uint32
	gl.GenBuffers(1, &EBO)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, EBO)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER,
		len(indices) * 4,
		gl.Ptr(indices),
		gl.STATIC_DRAW)

	return EBO
}

func main() {
	window := newWindow()

	defer glfw.Terminate()

	vertices := []float32{
		0.5, 0.5, 0.0,  // top right
		0.5, -0.5, 0.0, // bottom right
		-0.5, -0.5, 0.0, // bottom left
		-0.5, 0.5, 0.0, // top left
	}

	indices := []uint32 {
		0, 1, 3, // first triangle
		1, 2, 3, // second triangle
	}


	var VAO uint32
	gl.GenVertexArrays(1, &VAO)
	gl.BindVertexArray(VAO)

	bindVerts(vertices)

	EBO := createEBO(indices)

	shaderProg := triangleShaderProg()

	configureVAO(EBO)

	gl.PolygonMode(gl.FRONT_AND_BACK, gl.LINE)
	for !window.ShouldClose() {
		// Do OpenGL stuff.
		processInput(window)
		clearToCyberpunkColor()
		gl.UseProgram(shaderProg)
		gl.BindVertexArray(VAO)
		gl.DrawElements(gl.TRIANGLES, 6, gl.UNSIGNED_INT, gl.PtrOffset(0))
		gl.BindVertexArray(0)
		window.SwapBuffers()
		glfw.PollEvents()
	}
}
