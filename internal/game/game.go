package game

import (
	"image/color"

	"github.com/danmrichards/gba-pong/internal/display"
	"github.com/danmrichards/gba-pong/internal/input"
	"tinygo.org/x/drivers"
	"tinygo.org/x/tinydraw"
)

var (
	// Using a load of package level variables here because the compiler and
	// resulting ROM acts weirdly if using structs and receiver methods (e.g.
	// incorrect colours). Possibly something to do with alignment and memory
	// mapping.

	// Screen is the display where the game will draw pixels.
	Screen drivers.Displayer

	// Background is the background colour of the game.
	Background color.RGBA

	// PaddleWidth is the width of the player paddle.
	PaddleWidth = int16(10)

	// PaddleHeight is the height of the player paddle.
	PaddleHeight = int16(50)

	// PaddleColour is the colour of the player paddle.
	PaddleColour color.RGBA

	// BallWidth is the width of the ball.
	BallWidth = int16(8)

	// BallHeight is the height of the ball.
	BallHeight = int16(8)

	// BallColour is the colour of the ball.
	BallColour color.RGBA

	// paddle positional variables.
	paddleX = int16(10)
	paddleY = int16(10)

	// ball positional variables.
	ballX         = (display.Width / 2) - (BallWidth / 2)
	ballY         = (display.Height / 2) - (BallHeight / 2)
	ballXVelocity = int16(-2)
	ballYVelocity = int16(-1)
	ballMaxX      = display.Width - BallWidth
	ballMaxY      = display.Height - BallHeight
)

// Init initialises the game state.
func Init() {
	tinydraw.FilledRectangle(
		Screen, 0, 0, display.Width, display.Height, Background,
	)
}

// Update updates the current game state, intended to be called for each frame.
func Update() {
	clearPrevFrame()

	handleInput()

	updatePaddle()

	updateBall()
}

// clearPrevFrame clears the previous frame drawn pixels.
func clearPrevFrame() {
	// Clear previous paddle position.
	tinydraw.FilledRectangle(
		Screen,
		paddleX,
		paddleY,
		PaddleWidth,
		PaddleHeight,
		Background,
	)

	// Clear previous ball position.
	tinydraw.FilledRectangle(
		Screen,
		ballX,
		ballY,
		BallWidth,
		BallHeight,
		Background,
	)
}

// handleInput handles input from the keypad, updating game state accordingly.
func handleInput() {
	// Update paddle position.
	switch {
	case input.KeyPressed(input.KeyUp):
		paddleY -= 5
	case input.KeyPressed(input.KeyDown):
		paddleY += 5
	}
}

// updatePaddle updates the position of the paddle.
func updatePaddle() {
	// Stop the paddle from going off screen.
	paddleY = clamp(paddleY, 10, display.Height-PaddleHeight-10)

	// Draw paddle.
	tinydraw.FilledRectangle(
		Screen,
		10,
		paddleY,
		PaddleWidth,
		PaddleHeight,
		PaddleColour,
	)
}

// updateBall updates the position of the ball.
func updateBall() {
	collideEdge()
	collidePaddle()

	ballX = clamp(ballX+ballXVelocity, 0, ballMaxX)
	ballY = clamp(ballY+ballYVelocity, 0, ballMaxY)

	// Draw ball.
	tinydraw.FilledRectangle(
		Screen,
		ballX,
		ballY,
		BallWidth,
		BallHeight,
		BallColour,
	)
}

// collideEdge detects if the ball has collided with the edge of the screen
// and updates the direction accordingly.
func collideEdge() {
	if ballY == 0 || ballY == ballMaxY {
		ballYVelocity = -ballYVelocity
	} else if ballX == 0 || ballX == ballMaxX {
		ballXVelocity = -ballXVelocity
	}
}

// collidePaddle detects if the ball has collided with the paddle and updates
// the direction accordingly.
func collidePaddle() {
	if (ballX >= paddleX && ballX <= paddleX+PaddleWidth) &&
		(ballY >= paddleY && ballY <= paddleY+PaddleHeight) {
		ballX = paddleX + PaddleWidth
		ballXVelocity = -ballXVelocity
	}
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
