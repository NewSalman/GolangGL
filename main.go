package main

import (
	"runtime"
	"github.com/go-gl/glfw/v3.3/glfw"	
)

func Init() {
	runtime.LockOSThread()
}

func main() {
	err := glfw.Init()
	if err != nil {
		panic(err)
	}

	defer glfw.Terminate()

	window, err := glfw.CreateWindow(800, 600, "Mincraft Golang edition", nil, nil)

	window.MakeContextCurrent()

	for !window.ShouldClose() {
		// Do OpenGL stuff.
		window.SwapBuffers()
		glfw.PollEvents()
	}

}