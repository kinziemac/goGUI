package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

func main() {
	fmt.Println("hello go")
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Println("initializing SDL:", err)
		return
	}

	window, err := sdl.CreateWindow(
		"Game Start",
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		screenWidth,
		screenHeight,
		sdl.WINDOW_OPENGL)

	if err != nil {
		fmt.Println("initializing SDL:", err)
		return
	}

	// Destroys window at end of execution
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println("initializing SDL:", err)
		return
	}

	// Destroys renderer at end of execution
	defer renderer.Destroy()

	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {

			// For exiting window
			case *sdl.QuitEvent:
				return

			// State is for clicking (0 off: 1 on), Need x and y coordinates
			case *sdl.MouseMotionEvent:
				fmt.Printf("Mouse at: %+v\n ", event)
			}

		}

		renderer.SetDrawColor(255, 0, 0, 255)
		renderer.Clear()

		renderer.Present()

	}

}
