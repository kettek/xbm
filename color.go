package xbm

import (
	"image/color"
)

// BitColor represents a 1-bit color value.
type BitColor struct {
	Filled bool
}

// RGBA returns the RGBA value of the BitColor.
func (c BitColor) RGBA() (r, g, b, a uint32) {
	if c.Filled {
		return 0, 0, 0, 0xffff
	}
	return 0, 0, 0, 0
}

func toBitColor(c color.Color) color.Color {
	if _, ok := c.(BitColor); ok {
		return c
	}
	r, g, b, a := c.RGBA()

	return BitColor{
		uint8((0.2125*float32(r))+(0.7154*float32(g))+(0.0721*float32(b))/float32(255/a)) >= 32, // I guess?
	}
}

// BitColorModel is the ColorModel associated with the BitColor type.
var BitColorModel color.Model = color.ModelFunc(toBitColor)
