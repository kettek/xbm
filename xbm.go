package xbm

import (
	"image"
)

// Info contains additional XBM metadata.
type Info struct {
	// X,Y coordinate that determines the XBM hotspot.
	Hotspot image.Point
}
