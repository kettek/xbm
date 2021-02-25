package xbm

import (
	"image"
	"image/color"
)

type Bits struct {
	// Pix holds the image's pixels, as boolean values.
	Pix []bool
	// Stride is the Pix stride (in bytes) between vertically adjacent pixels.
	Stride int
	// Rect is the image's bounds.
	Rect image.Rectangle
}

func (b *Bits) ColorModel() color.Model { return BitColorModel }

func (b *Bits) Bounds() image.Rectangle { return b.Rect }

func (b *Bits) At(x, y int) color.Color {
	if !(image.Point{x, y}.In(b.Rect)) {
		return BitColor(false)
	}
	i := b.PixOffset(x, y)
	return BitColor(b.Pix[i])
}

func (b *Bits) PixOffset(x, y int) int {
	return (y-b.Rect.Min.Y)*b.Stride + (x - b.Rect.Min.X)
}

func (b *Bits) Set(x, y int, c color.Color) {
	if !(image.Point{x, y}.In(b.Rect)) {
		return
	}
	i := b.PixOffset(x, y)
	b.Pix[i] = bool(BitColorModel.Convert(c).(BitColor))
}

func (b *Bits) SetBit(x, y int, c bool) {
	if !(image.Point{x, y}.In(b.Rect)) {
		return
	}
	i := b.PixOffset(x, y)
	b.Pix[i] = c
}

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

func (b *Bits) Opaque() bool {
	return true
}

func NewBits(r image.Rectangle) *Bits {
	return &Bits{
		Pix:    make([]bool, r.Max.Y*r.Max.X+r.Max.X),
		Stride: r.Max.X,
		Rect:   r,
	}
}
