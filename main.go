package main

import (
	"unsafe"

	"runtime/volatile"
)

const (
	screenWidth  = 240
	screenHeight = 160

	memIO   = 0x04000000 // I/O Registers.
	memPAL  = 0x05000000 // Colour Palette RAM.
	memVRAM = 0x06000000 // Video RAM.
	memOAM  = 0x07000000 // Object Attribute Memory (Sprite hardware).

	// Keys.
	keyUP   = 0x0040
	keyDown = 0x0080
	keyAny  = 0x03FF

	objectAttr0YMask = 0x0FF
	objectAttr1XMask = 0x1FF
)

var (
	// display represents the I/O register for LCD control.
	//
	// Used to set video modes.
	display = (*volatile.Register16)(unsafe.Pointer(uintptr(memIO)))

	// vcount represents the vertical counter register.
	//
	// Holds the index of the current row being drawn to screen. Can be used
	// to detect when the screen has been totally updated (VBLANK).
	vcount = (*volatile.Register16)(unsafe.Pointer(uintptr(memIO + 0x0006)))

	// keyPad represents the GBA key pad.
	keyPad = (*volatile.Register16)(unsafe.Pointer(uintptr(memIO + 0x0130)))

	// oam represents the object attribute memory (OAM).
	//
	// The OAM is onboard hardware used for rendering sprites.
	oam = (*objectAttributes)(unsafe.Pointer(uintptr(memOAM)))

	// tileMem represents a set of 16kb blocks of video RAM used to store tiles.
	tileMem = (*tileBlock)(unsafe.Pointer(uintptr(memVRAM)))

	// objectPaletteMem represents the colour palette memory used for objects.
	objectPaletteMem = (*rgb15)(unsafe.Pointer(uintptr(memPAL + 0x200)))
)

type (
	// rgb15 represents a 15-bit colours GBA colour.
	//
	// There are 15 actual bits, but colours are aligned as 16-bit.
	//
	// ?BBBBBGGGGGRRRRR - One unused bit, 5 blue, 5 green, 5 red.
	rgb15 uint16

	// objectAttributes represents the attributes of a sprite in the GBA
	// Object Access Memory (OAM).
	objectAttributes struct {
		attr0 uint16 // Y coordinate, shape and colour mode of the object's tiles (4bpp or 8bpp).
		attr1 uint16 // X coordinate and size.
		attr2 uint16 // Base tile index and the colour palette (when in 4bpp mode).
		pad   uint16
	}

	// tile4BPP represents a 4-bit per pixel tile.
	//
	// 8 x 8 at 4-bits per pixel = 8 ints
	//
	// Each tile can have 16 colours (i.e. reference 16 colour palette values).
	tile4BPP [8]uint32

	// tileBlock represents a 16KB section of vRAM used for tiles.
	//
	// 512 4-bit tiles can fit in the block.
	tileBlock [512]tile4BPP
)

func main() {

}
