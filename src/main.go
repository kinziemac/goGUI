package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenWidth  = 700
	screenHeight = 700
	boxSize      = 1
	totalScreen  = screenWidth * screenHeight
	rectSize     = screenHeight
)

func main() {
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

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println("initializing SDL:", err)
		return
	}

	// Destroys window and renderer
	defer window.Destroy()
	defer renderer.Destroy()

	rectArray := createRectArray(
		screenWidth,
		screenHeight,
		totalScreen)

	b1 := initBlock(1, 0, 0, 700)
	b2 := initBlock(1, 100, 0, 700)

	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch val := event.(type) {

			case *sdl.QuitEvent:
				return

			case *sdl.KeyboardEvent:
				if val.Keysym.Sym == sdl.K_SPACE {
					fmt.Println("SPACE pressed")

					b1.drawBlock(renderer)
					b2.drawBlock(renderer)

					// for i := 0; i < len(b.pixels); i++ {
					// 	renderer.SetDrawColor(255, 255, 255, 255)
					// 	renderer.FillRect(&b.pixels[i])
					// }
				}
			}
		}

		//checks if mouse is clicked
		mouseX, mouseY, mouseButtonState := sdl.GetMouseState()

		if mouseButtonState == 1 {
			fmt.Printf("Mouse at x: %+v, y: %+v, state: %+v\n", mouseX, mouseY, mouseButtonState)
			index := int((mouseX) + rectSize*(mouseY))

			//from player.go
			colorRect(renderer, &rectArray[index])
		}

		renderer.Present()
	}

}
