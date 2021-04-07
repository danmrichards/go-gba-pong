# GO GBA Pong [![License](http://img.shields.io/badge/license-mit-blue.svg)](https://raw.githubusercontent.com/danmrichards/go-gba-pong/main/LICENSE) [![Go Report Card](https://goreportcard.com/badge/github.com/danmrichards/go-gba-pong)](https://goreportcard.com/report/github.com/danmrichards/go-gba-pong)
A basic pong-like game for the Game Boy Advance in [TinyGo][1].

##  Requirements
* Go 1.13+
* TinyGo
* [GBA Fix][2]. This is required to pad and patch header information onto the compiled ROM file. TinyGo does not ship with the header information in its GBA target.

## Building From Source
```bash
$ make build
```

## Usage
After building from source you should see a ROM file called `pong.GBA` in the `bin` directory.

This file can be loaded into a Game Boy Advance emulator or onto a physical device if you have a flash cartridge.

![emulator][screenshot]

[1]: https://tinygo.org/
[2]: https://github.com/devkitPro/gba-tools

[screenshot]: screenshot.png "Screenshot"
