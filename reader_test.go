package xbm

import (
	"fmt"
	"image"
	"os"
	"testing"
)

func TestReadAndPrint(t *testing.T) {
	info, m, err := DecodeXBM("tests/goblin.xbm")
	if err != nil {
		t.Error(err)
	}
	b := m.Bounds()
	fmt.Printf("Decoded XBM Dimensions: %dx%d->%dx%d\n", b.Min.X, b.Min.Y, b.Max.X, b.Max.Y)
	fmt.Printf("Metadata: data=%s width=%s height=%s hotspot=%d,%d hotspotX=%s hotspotY=%s\n", info.DataName, info.WidthName, info.HeightName, info.Hotspot.X, info.Hotspot.Y, info.HotspotXName, info.HotspotYName)
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

func DecodeXBM(path string) (Info, image.Image, error) {
	f, err := os.Open(path)
	if err != nil {
		return Info{}, nil, err
	}
	defer f.Close()

	info, img, err := DecodeInfoAndImage(f)
	if err != nil {
		return info, nil, err
	}

	return info, img, err
}
