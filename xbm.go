package xbm

import (
	"image"
)

// Info contains additional XBM metadata.
type Info struct {
	// X,Y coordinate that determines the XBM hotspot.
	Hotspot image.Point
}

// XBM contains a general XBM format.
type XBM struct {
	Info
	image.Image
}
