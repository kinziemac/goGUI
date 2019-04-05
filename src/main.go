package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenWidth      = 600
	screenHeight     = 600
	boxSize          = 10
	totalScreen      = screenWidth * screenHeight
	screenDimensions = totalScreen / boxSize
	rectSize         = screenHeight / boxSize
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

	rectArray := make([]sdl.Rect, screenDimensions)
	xCoor := 0
	yCoor := 0
	for i := 0; i < screenDimensions; i++ {
		rectArray[i] = sdl.Rect{int32(xCoor), int32(yCoor), boxSize, boxSize}
		xCoor = xCoor + boxSize

		// Basically moving left to right and then reset a row down
		if xCoor >= screenWidth {
			xCoor = 0
			yCoor = yCoor + boxSize
		}
	}

	// Destroys renderer at end of execution
	defer renderer.Destroy()

	// startUp := false
	prevX := int32(0)
	prevY := int32(0)

	for {
		//initialize gameboard
		// if startUp == false {
		// 	renderer.SetDrawColor(255, 255, 255, 255)
		// 	renderer.Clear()

		// 	for i := 0; i < len(rectArray); i++ {
		// 		renderer.SetDrawColor(0, 0, 0, 255)
		// 		renderer.DrawRect(&rectArray[i])
		// 	}

		// 	renderer.Present()
		// 	fmt.Println("turned screen white")
		// 	startUp = true
		// }
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch val := event.(type) {
			// For exiting window
			case *sdl.QuitEvent:
				return

			case *sdl.KeyboardEvent:

				// Don't need to press this anymore but just if I want to use the keydown thing
				if val.Keysym.Sym == sdl.K_ESCAPE {
					fmt.Println("ESC pressed")

					for i := 0; i < len(rectArray); i++ {
						renderer.SetDrawColor(0, 0, 0, 255)
						renderer.DrawRect(&rectArray[i])
					}
				}
			}
		}

		//checks if mouse is clicked
		mouseX, mouseY, mouseButtonState := sdl.GetMouseState()

		if mouseX != prevX && mouseY != prevY && mouseButtonState == 1 {
			fmt.Printf("Mouse at x: %+v, y: %+v, state: %+v\n", mouseX, mouseY, mouseButtonState)
			prevX = mouseX
			prevY = mouseY

			index := int((mouseX / boxSize) + rectSize*(mouseY/boxSize))

			if index >= (rectSize)*(rectSize) {
				fmt.Println("out of index:")
				return
			}

			renderer.SetDrawColor(0, 0, 255, 255)
			renderer.FillRect(&rectArray[index])
			// renderer.Present()
		}

		renderer.Present()
		renderer.SetDrawColor(255, 255, 255, 255)
		renderer.Clear()

	}

}
