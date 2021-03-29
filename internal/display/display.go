package display

import (
	"image/color"
	"unsafe"

	"runtime/volatile"

	"tinygo.org/x/drivers"
	"tinygo.org/x/tinydraw"
)

const (
	// Width is the width of the GBA screen.
	Width = 240

	// Height is the height of the GBA screen.
	Height = 160
)

// vcount represents the vertical counter register.
//
// Holds the index of the current row being drawn to screen. Can be used
// to detect when the screen has been totally updated (VBLANK).
var vcount = (*volatile.Register16)(unsafe.Pointer(uintptr(0x04000006)))

// Clear clears the screen to the given colour.
func Clear(d drivers.Displayer, c color.RGBA) {
	tinydraw.FilledRectangle(d, 0, 0, Width, Height, c)
}

// VSync is a poor mans vertical sync. It blocks until a vertical blank (VBLANK)
// is detected from the VCOUNT register.
//
// A VBLANK indicates that all the rows on the screen have been updated. The
// hardware will pause after a VBLANK, in order to avoid screen tearing updates
// to the screen buffer should occur during the pause.
//
// This is highly inefficient due to the busy waiting, it churns CPU and wastes
// battery. But it's the best we can do in Go due to patchy support for the
// interrupt API.
func VSync() {
	for vcount.Get() >= 160 {
	}

	for vcount.Get() < 160 {
	}
}
