package xbm

import (
	"fmt"
	"image"
	"image/color"
	"os"
	"testing"
)

func TestWrite(t *testing.T) {
	encoder := Encoder{}
	f, err := os.Create("tests/write.xbm")
	if err != nil {
		t.Error(err)
	}
	m := image.NewRGBA(image.Rect(0, 0, 7, 8))
	for i := 0; i < 8; i++ {
		m.Set(i, 4, color.RGBA{255, 255, 255, 255})
	}
	encoder.Encode(f, m)
}

func TestWriteRead(t *testing.T) {
	encoder := Encoder{}
	f, err := os.Create("tests/write.xbm")
	if err != nil {
		t.Error(err)
	}
	m := image.NewRGBA(image.Rect(0, 0, 4, 4))
	m.Set(0, 1, color.RGBA{64, 0, 0, 255})
	m.Set(1, 1, color.RGBA{64, 0, 0, 255})
	m.Set(2, 1, color.RGBA{0, 64, 0, 255})
	m.Set(3, 1, color.RGBA{0, 64, 0, 255})
	encoder.Encode(f, m)

	r, err := os.Open("tests/write.xbm")
	if err != nil {
		t.Error(err)
	}

	info, image, err := DecodeInfoAndImage(r)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%+v, %+v\n", info, image)

}
