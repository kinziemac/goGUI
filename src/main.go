package main

import (
	"fmt"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenDim     = 700
	blockDim      = 100
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
					fmt.Println("SPACE pressed - this creates the board, may need to hit a couple times")

					//creates board
					for i := 0; i < len(blockArray); i++ {
						blockArray[i].renderBlock(renderer)
					}
					renderer.Present()
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

			//if block is currently unfinished or not owned by anyone and if the user is currently writing on it
			if blockArray[boxIndex].isAllowed(&p) {
				p.active = true
				blockArray[boxIndex].drawOnBlock(renderer, int(mouseX), int(mouseY), blockDim, &p)
				renderer.Present()

			} else if p.currentBlock != boxIndex {
				//do nothing, they've gone out of the lines

			} else {
				p.currentBlock = -1
			}

			//when player unclicks
		} else {
			if p.active {
				if blockArray[p.currentBlock].blockFilled() {
					blockArray[p.currentBlock].completeBlock(&p, renderer)
					fmt.Println("You coloured all of it!")
					time.Sleep(20 * time.Millisecond)

				} else {
					blockArray[p.currentBlock].resetBlock(renderer)
					fmt.Println("You didn't colour all of it :(")
				}

				p.currentBlock = -1
				p.active = false
				renderer.Present()
			}

		}

		// renderer.Present()
	}

}
