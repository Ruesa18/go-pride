package main

import (
	"fmt"

	"github.com/gookit/color"
)

func displayFlag(flagWidth int, colors []uint8) {
	for i := 0; i < len(colors); i++ {
		for x := 0; x < flagWidth; x++ {
			color.S256(colors[i], colors[i]).Print(" ")
		}
		fmt.Println()
	}
}

func main() {
	flagWidth := 20
	colors := []uint8{160, 166, 220, 28, 4, 90}

	displayFlag(flagWidth, colors)
}
