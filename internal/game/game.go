package game

import (
	"image/color"

	"runtime/interrupt"

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

	// Paddle positional variables.
	paddleX = int16(10)
	paddleY = int16(10)

	// Ball positional variables.
	ballX             int16
	ballY             int16
	ballXVelocity     int16
	ballDown          = int16(1)
	ballUp            = int16(-1)
	ballYVelocity     = ballUp
	ballLastYVelocity int16
	ballMaxX          = display.Width - BallWidth
	ballMaxY          = display.Height - BallHeight

	// Paddle miss indicator and frame timer. The time package does not work
	// well in tinygo, so instead we just wait for a number of frames.
	miss       bool
	missFrames int
)

// Init initialises the game state.
func Init() {
	display.Clear(Screen, Background)

	resetBall(true, true)
}

// Update updates the current game state, intended to be called for each frame.
func Update(interrupt.Interrupt) {
	clearPrevFrame()

	// The last shot missed the paddle, reset the ball position and wait for
	// the remaining frames before starting again.
	if miss {
		if missFrames > 0 {
			missFrames--
		} else {
			miss = false
			resetBall(true, true)
		}
	}

	handleInput()

	updatePaddle()

	updateBall()
}

// resetBall resets the position and/or velocity of the ball to the default
// values.
func resetBall(position, velocity bool) {
	if position {
		ballX = (display.Width / 2) - (BallWidth / 2)
		ballY = (display.Height / 2) - (BallHeight / 2)
	}

	if velocity {
		ballXVelocity = int16(-2)

		// Alternate ball direction on each new shot.
		if ballLastYVelocity == ballDown {
			ballYVelocity = ballUp
			ballLastYVelocity = ballUp
		} else {
			ballYVelocity = ballDown
			ballLastYVelocity = ballDown
		}
	}
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

// collideEdge detects if the ball has collided with the edge of the screen.
func collideEdge() {
	if ballY == 0 || ballY == ballMaxY {
		// Top or bottom of screen bounce.
		ballYVelocity = -ballYVelocity
	} else if ballX == ballMaxX {
		// Right hand side bounce.
		ballXVelocity = -ballXVelocity
	} else if ballX == 0 {
		// Paddle missed, stop the ball and set the miss frame counter.
		ballXVelocity = 0
		ballYVelocity = 0
		resetBall(true, false)
		missFrames = 10
		miss = true
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
