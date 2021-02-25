package xbm

import (
	"fmt"
	"image"
	"os"
	"testing"
)

func TestReadAndPrint(t *testing.T) {
	m, err := DecodeXBM("tests/goblin.xbm")
	if err != nil {
		t.Error(err)
	}
	b := m.Bounds()
	fmt.Printf("Decoded XBM Dimensions: %dx%d->%dx%d\n", b.Min.X, b.Min.Y, b.Max.X, b.Max.Y)
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			v, ok := m.At(x, y).(BitColor)
			if !ok {
				continue
			}
			if v {
				fmt.Printf(".")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Println("")
	}
}

func DecodeXBM(path string) (image.Image, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	img, err := Decode(f)
	if err != nil {
		return nil, err
	}

	return img, err
}
