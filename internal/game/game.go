package game

import (
	"image/color"

	"github.com/danmrichards/gba-pong/internal/display"
	"github.com/danmrichards/gba-pong/internal/input"
	"tinygo.org/x/drivers"
	"tinygo.org/x/tinydraw"
)

const (
	paddleWidth  = 10
	paddleHeight = 50
)

// Pong represents the game state.
type Pong struct {
	screen drivers.Displayer
	bg     color.RGBA

	paddleY      int16
	paddleColour color.RGBA
}

// NewPong returns an instantiated game state configured to draw to the given
// display in the given colours.
func NewPong(screen drivers.Displayer, bg, paddleColour color.RGBA) *Pong {
	return &Pong{
		screen:       screen,
		bg:           bg,
		paddleY:      10,
		paddleColour: paddleColour,
	}
}

// Update updates the current game state, intended to be called for each frame.
func (p *Pong) Update() {
	// Clear previous paddle position.
	tinydraw.FilledRectangle(
		p.screen,
		10,
		p.paddleY,
		paddleWidth,
		paddleHeight,
		p.bg,
	)

	// Update paddle position.
	switch {
	case input.KeyPressed(input.KeyUp):
		p.paddleY -= 5
	case input.KeyPressed(input.KeyDown):
		p.paddleY += 5
	}

	p.paddleY = clamp(p.paddleY, 10, display.Height-paddleHeight-10)

	// Draw new paddle.
	tinydraw.FilledRectangle(
		p.screen,
		10,
		p.paddleY,
		paddleWidth,
		paddleHeight,
		p.paddleColour,
	)
}

// clamp returns n normalised to the range min <= n <= max.
func clamp(n, min, max int16) int16 {
	if n < min {
		return min
	} else if n > max {
		return max
	}

	return n
}
