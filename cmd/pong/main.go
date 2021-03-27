package main

import (
	"image/color"

	"machine"

	"github.com/danmrichards/gba-pong/internal/display"
	"github.com/danmrichards/gba-pong/internal/game"
)

var (
	screen = machine.Display

	white = color.RGBA{
		R: 255,
		G: 255,
		B: 255,
		A: 255,
	}

	black = color.RGBA{}

	paddleY = int16(10)
)

func main() {
	// Enable bitmap mode.
	screen.Configure()

	display.Clear(screen, black)

	p := game.NewPong(screen, black, white)

	for {
		display.VSync()

		// Update the game state.
		p.Update()
	}
}
