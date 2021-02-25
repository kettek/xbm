package xbm

import (
	"image"
)

// XBMInfo contains additional XBM metadata.
type XBMInfo struct {
	// X,Y coordinate that determines the XBM hotspot.
	Hotspot image.Point
}
