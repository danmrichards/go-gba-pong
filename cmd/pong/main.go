package main

import (
	"image/color"

	"machine"
	"runtime/interrupt"

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
)

func main() {
	// Enable bitmap mode.
	screen.Configure()

	display.VSync()

	// Configure the game.
	game.Screen = screen
	game.Background = black
	game.PaddleColour = white
	game.BallColour = white

	game.Init()

	// Register an interrupt to only update the game state when the screen is
	// done drawing.
	interrupt.New(machine.IRQ_VBLANK, game.Update).Enable()

	// Infinite loop is required to prevent the ROM from exiting.
	for {
	}
}
