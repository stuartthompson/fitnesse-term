// Copyright 2018 Stuart Thompson.

// This file is part of Fitnesse-Term.

// Fitnesse-Term is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// Fitnesse-Term is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with Fitnesse-Term. If not, see <http://www.gnu.org/licenses/>.

package io

import (
	"unicode/utf8"

	"github.com/nsf/termbox-go"
)

// Characters for rendering the border runes
const (
	BorderRuneHorizontal        = rune('\u2550')
	BorderRuneVertical          = rune('\u2551')
	BorderRuneCornerTopLeft     = rune('\u2554')
	BorderRuneCornerTopRight    = rune('\u2557')
	BorderRuneCornerBottomLeft  = rune('\u255A')
	BorderRuneCornerBottomRight = rune('\u255D')
)

// ClearArea ...
// Clears an area to the specified background color.
func ClearArea(x int, y int, width int, height int, bgColor int) {
	for ix := 0; ix < width; ix++ {
		for iy := 0; iy < height; iy++ {
			termbox.SetCell(x+ix, y+iy, ' ', 0, termbox.Attribute(bgColor))
		}
	}
}

// ClearScreen ...
// Clears the screen using a specified color.
func ClearScreen(bgColor int) {
	termbox.Clear(0, termbox.Attribute(bgColor))
}

// Flush ...
// Flush render commands.
func Flush() {
	termbox.Flush()
}

// GetWindowSize ...
// Gets the current dimensions of the terminal.
func GetWindowSize() (width int, height int) {
	return termbox.Size()
}

// RenderPaneBorder ...
// Renders a border for a window pane.
func RenderPaneBorder(x int, y int, width int, height int, fgColor int, bgColor int) {
	// Render the corners
	termbox.SetCell(x, y, BorderRuneCornerTopLeft, termbox.Attribute(fgColor), termbox.Attribute(bgColor))
	termbox.SetCell(x+width, y, BorderRuneCornerTopRight, termbox.Attribute(fgColor), termbox.Attribute(bgColor))
	termbox.SetCell(x, y+height, BorderRuneCornerBottomLeft, termbox.Attribute(fgColor), termbox.Attribute(bgColor))
	termbox.SetCell(x+width, y+height, BorderRuneCornerBottomRight, termbox.Attribute(fgColor), termbox.Attribute(bgColor))
	// Render top border
	for ix := 1; ix < width; ix++ {
		termbox.SetCell(x+ix, y, BorderRuneHorizontal, termbox.Attribute(fgColor), termbox.Attribute(bgColor))
	}
	// Render bottom border
	for ix := 1; ix < width; ix++ {
		termbox.SetCell(x+ix, y+height, BorderRuneHorizontal, termbox.Attribute(fgColor), termbox.Attribute(bgColor))
	}
	// Render left border
	for iy := 1; iy < height; iy++ {
		termbox.SetCell(x, y+iy, BorderRuneVertical, termbox.Attribute(fgColor), termbox.Attribute(bgColor))
	}
	// Render right border
	for iy := 1; iy < height; iy++ {
		termbox.SetCell(x+width, y+iy, BorderRuneVertical, termbox.Attribute(fgColor), termbox.Attribute(bgColor))
	}
}

// RenderText ...
// Renders a string at specific coordinates using the supplied colors.
func RenderText(text string, x int, y int, fgColor int, bgColor int) {
	// Get the text to render as a byte array
	bytes := []byte(text)

	// Initialize terminal column index
	colIx := x

	// Loop through the bytes in the slice
	for bIx := 0; bIx < len(bytes); {
		// Decode the next rune in the string
		r, size := utf8.DecodeRune(bytes[bIx:])
		// Set the cell value to that of the decoded rune
		termbox.SetCell(colIx, y, r, termbox.Attribute(fgColor), termbox.Attribute(bgColor))
		// Advance the index by the size of the rune just decoded (some runes use multiple bytes)
		bIx += size
		// Increment the terminal column index (one column per rune rendered)
		colIx++
	}
}
