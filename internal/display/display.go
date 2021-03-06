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

	// Interrupt request flags for VBLANK and HBLANK.
	vblankIRQ = 1 << 3
	hblankIRQ = 1 << 4
)

var (
	// vcount represents the vertical counter register.
	//
	// Holds the index of the current row being drawn to screen. Can be used
	// to detect when the screen has been totally updated (VBLANK).
	vcount = (*volatile.Register16)(unsafe.Pointer(uintptr(0x04000006)))

	// displayStatus represents the display status and interrupt control.
	displayStatus = (*volatile.Register32)(unsafe.Pointer(uintptr(0x4000004)))
)

// Clear clears the screen to the given colour.
func Clear(d drivers.Displayer, c color.RGBA) {
	tinydraw.FilledRectangle(d, 0, 0, Width, Height, c)
}

// VSync enables the interrupts for VBLANK and HBLANK, effectively enabling
// a vertical sync mode.
func VSync() {
	displayStatus.SetBits(vblankIRQ | hblankIRQ)
}

// PoorVSync is a poor mans vertical sync. It blocks until a vertical blank
// (VBLANK) is detected from the VCOUNT register.
//
// A VBLANK indicates that all the rows on the screen have been updated. The
// hardware will pause after a VBLANK, in order to avoid screen tearing updates
// to the screen buffer should occur during the pause.
//
// This is highly inefficient due to the busy waiting, it churns CPU and wastes
// battery.
func PoorVSync() {
	for vcount.Get() >= 160 {
	}

	for vcount.Get() < 160 {
	}
}
