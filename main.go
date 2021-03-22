package main

import (
	"image/color"
	"unsafe"

	"machine"
	"runtime/interrupt"
	"runtime/volatile"
)

const (
	screenWidth  = 240
	screenHeight = 160
)

var (
	// displayCtl represents the I/O register for LCD control.
	//
	// Used to set video modes.
	displayCtl = (*volatile.Register32)(unsafe.Pointer(uintptr(0x04000000)))

	// displayStatus represents the display status and interrupt control.
	displayStatus = (*volatile.Register32)(unsafe.Pointer(uintptr(0x4000004)))

	screen = machine.Display

	white = color.RGBA{
		R: 255,
		G: 255,
		B: 255,
		A: 255,
	}

	black = color.RGBA{}

	rectX = int16(0)
)

func update(interrupt.Interrupt) {
	// Wrap to edge of screen.
	if rectX > screenWidth*screenHeight/10 {
		rectX = 0
	} else if rectX > 0 {
		// Reset the previous rectangle to black. Gives the illusion of
		// the rectangle moving.
		last := rectX - 10
		drawRectangle(last%screenWidth, (last/screenWidth)*10, 10, 10, black)
	}

	// Draw a white rectangle.
	drawRectangle(rectX%screenWidth, (rectX/screenWidth)*10, 10, 10, white)

	rectX += 10
}

// drawRectangle sets a number of pixels to a given color in order to form a
// rectangle using the given co-ordinates, width and height.
func drawRectangle(x, y, w, h int16, c color.RGBA) {
	for ry := y; ry < y+w; ry++ {
		for rx := x; rx < x+h; rx++ {
			screen.SetPixel(rx, ry, c)
		}
	}
}

// clear clears the screen to black.
func clear() {
	for x := int16(0); x < screenWidth; x++ {
		for y := int16(0); y < screenHeight; y++ {
			screen.SetPixel(x, y, black)
		}
	}
}

func main() {
	// Enable bitmap mode.
	screen.Configure()

	// Enable the interrupts for VBLANK and HBLANK.
	displayStatus.SetBits(1<<3 | 1<<4)

	clear()

	// Update the screen on VBLANK to avoid screen tearing.
	interrupt.New(machine.IRQ_VBLANK, update).Enable()

	for {
	}
}
