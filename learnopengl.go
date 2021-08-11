package main // import "github.com/go-gl/example/gl41core-cube"

import (
	"fmt"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func main() {
	err := glfw.Init()
	if err != nil {
		panic(err)
	}
	defer glfw.Terminate()

	gl.Init()

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
		fmt.Println("resize")
		gl.Viewport(0, 0, int32(width), int32(height))
	})

	for !window.ShouldClose() {
		// Do OpenGL stuff.
		window.SwapBuffers()
		glfw.PollEvents()
	}
}
