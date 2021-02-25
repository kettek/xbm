package xbm

import (
	"image"
)

// Info contains additional XBM metadata.
type Info struct {
	// X, Y coordinate that determines the XBM hotspot.
	Hotspot image.Point
	// Variable name used for the pixel data.
	DataName string
	// Variables names for the width and height.
	WidthName, HeightName string
	// Variable names for the x and y hotspot.
	HotspotXName, HotspotYName string
}
