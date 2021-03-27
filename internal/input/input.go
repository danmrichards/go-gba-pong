package input

import (
	"unsafe"

	"runtime/volatile"
)

// KeyCode represents a GBA keypad code.
type KeyCode uint16

var (
	// input represents the register for keypad input.
	//
	// The GBA is a bit strange and has "inverted" values for input, 0 = pressed
	// and 1 = unpressed.
	input = (*volatile.Register16)(unsafe.Pointer(uintptr(0x04000130)))

	// Keys.
	KeyA      = KeyCode(0x0001)
	KeyB      = KeyCode(0x0002)
	KeySelect = KeyCode(0x0004)
	KeyStart  = KeyCode(0x0008)
	KeyRight  = KeyCode(0x0010)
	KeyLeft   = KeyCode(0x0020)
	KeyUp     = KeyCode(0x0040)
	KeyDown   = KeyCode(0x0080)
	KeyR      = KeyCode(0x0100)
	KeyL      = KeyCode(0x0200)

	keyMask = uint16(0x03FF)
)

// KeyPressed returns true if the given key is pressed.
func KeyPressed(k KeyCode) bool {
	return (^input.Get()&keyMask)&uint16(k) > 0
}
