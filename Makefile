TARGET=gameboy-advance
ROM=pong.gba

.PHONY: build
build:
	tinygo build -o bin/${ROM} -target ${TARGET} main.go
	gbafix bin/${ROM}