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

	// blockArray := createBlockArray(
	blockArray := createBlockArray(
		screenDim,
		totalScreen,
		blockDim)

	// Create New Player
	p := newPlayer()

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

				if val.Keysym.Sym == sdl.K_ESCAPE {
					fmt.Println("Score is: ", p.score)
				}
			}
		}

		mouseX, mouseY, mouseButtonState := sdl.GetMouseState()
		if mouseButtonState == 1 {
			boxIndex := (mouseX / blockDim) + (mouseY/blockDim)*blocksPerPage

			//if the user has not touched a block yet
			if p.currentBlock == -1 {
				p.currentBlock = boxIndex
			}

			if blockArray[boxIndex].isAllowed() && p.currentBlock == boxIndex {
				p.active = true
				blockArray[boxIndex].drawOnBlock(renderer, int(mouseX), int(mouseY), blockDim)
			} else if p.currentBlock != boxIndex {
				//do nothing

			} else {
				p.currentBlock = -1
			}

		} else {
			if p.active {
				if blockArray[p.currentBlock].blockFilled() {
					blockArray[p.currentBlock].fillBlock(255, 0, 0, renderer)
					p.score++
					fmt.Println("You coloured all of it!")

				} else {
					blockArray[p.currentBlock].fillBlock(0, 0, 0, renderer)
					fmt.Println("You didn't colour all of it :(")
				}

				p.currentBlock = -1
				p.active = false
			}

		}

		renderer.Present()
	}

}
