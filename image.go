package xbm

import (
	"image"
	"image/color"
)

// Bits is an in-memory image whose At method returns color.BitColor values.
type Bits struct {
	// Pix holds the image's pixels, as boolean values.
	Pix []bool
	// Stride is the Pix stride (in bytes) between vertically adjacent pixels.
	Stride int
	// Rect is the image's bounds.
	Rect image.Rectangle
}

// ColorModel returns BitColorModel.
func (b *Bits) ColorModel() color.Model { return BitColorModel }

// Bounds returns the image's bounds.
func (b *Bits) Bounds() image.Rectangle { return b.Rect }

// At returns the pixel color at (x, y).
func (b *Bits) At(x, y int) color.Color {
	if !(image.Point{x, y}.In(b.Rect)) {
		return BitColor(false)
	}
	i := b.PixOffset(x, y)
	return BitColor(b.Pix[i])
}

// PixOffset returns the index of the first element of Pix that corresponds to
// the pixel at (x, y).
func (b *Bits) PixOffset(x, y int) int {
	return (y-b.Rect.Min.Y)*b.Stride + (x - b.Rect.Min.X)
}

// Set sets the binary pixel value at (x, y) to match the provided color,
// using BitColorModel's conversion algorithm.
func (b *Bits) Set(x, y int, c color.Color) {
	if !(image.Point{x, y}.In(b.Rect)) {
		return
	}
	i := b.PixOffset(x, y)
	b.Pix[i] = bool(BitColorModel.Convert(c).(BitColor))
}

// SetBit sets the binary pixel value at (x, y) to a boolean value.
func (b *Bits) SetBit(x, y int, c bool) {
	if !(image.Point{x, y}.In(b.Rect)) {
		return
	}
	i := b.PixOffset(x, y)
	b.Pix[i] = c
}

// SubImage returns an image representing the portion of the image p visible
// through r. The returned value shares pixels with the original image.
func (b *Bits) SubImage(r image.Rectangle) image.Image {
	r = r.Intersect(b.Rect)
	if r.Empty() {
		return &Bits{}
	}
	i := b.PixOffset(r.Min.X, r.Min.Y)
	return &Bits{
		Pix:  b.Pix[i:],
		Rect: r,
	}
}

// Opaque scans the entire image and reports whether it is fully opaque.
// This is always true for XBM images.
func (b *Bits) Opaque() bool {
	return true
}

// NewBits returns a new Bits image with the given bounds.
func NewBits(r image.Rectangle) *Bits {
	return &Bits{
		Pix:    make([]bool, r.Max.Y*r.Max.X+r.Max.X),
		Stride: r.Max.X,
		Rect:   r,
	}
}
