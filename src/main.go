package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenDim     = 700
	blockDim      = 50
	totalScreen   = screenDim * screenDim
	blocksPerPage = screenDim / blockDim
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
		screenDim,
		screenDim,
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

	// rectArray := createRectArray(
	// 	screenDim,
	// 	totalScreen)

	// blockArray := createBlockArray(
	blockArray := createBlockArray(
		screenDim,
		totalScreen,
		blockDim)

	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch val := event.(type) {

			case *sdl.QuitEvent:
				return

			case *sdl.KeyboardEvent:
				if val.Keysym.Sym == sdl.K_SPACE {
					fmt.Println("SPACE pressed")

					//creates board
					for i := 0; i < len(blockArray); i++ {
						blockArray[i].renderBlock(renderer)
					}
				}
			}
		}

		mouseX, mouseY, mouseButtonState := sdl.GetMouseState()

		if mouseButtonState == 1 {

			boxIndex := (mouseX / blockDim) + (mouseY/blockDim)*blocksPerPage

			//from player.go
			blockArray[boxIndex].drawOnBlock(renderer, int(mouseX), int(mouseY), blockDim)
			// colorRect(renderer, &rectArray[index])
			// colorRect(renderer, &rectArray[index])
		}

		renderer.Present()
	}

}
