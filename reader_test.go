package xbm

import (
	"fmt"
	"image"
	"os"
	"testing"
)

func TestReadAndPrint(t *testing.T) {
	info, m, err := DecodeXBM("tests/helmet.xbm")
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

func ExampleDecode() {
	n := "beholder.xbm"
	f, err := os.Open(n)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	m, err := Decode(f)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Read XBM: %s, size: %dx%d\n", n, m.Bounds().Max.X, m.Bounds().Max.Y)
}

func ExampleDecode_ascii() {
	n := "beholder.xbm"
	f, err := os.Open(n)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	m, err := Decode(f)
	if err != nil {
		panic(err)
	}
	b := m.Bounds()

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

func ExampleDecodeConfig() {
	n := "beholder.xbm"
	f, err := os.Open(n)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	c, err := DecodeConfig(f)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Read XBM: %s, size: %dx%d\n", n, c.Width, c.Height)
}

func ExampleDecodeInfo() {
	n := "cursor.xbm"
	f, err := os.Open(n)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	info, err := DecodeInfo(f)
	if err != nil {
		panic(err)
	}
	fmt.Printf("XBM Metadata: data=%s width=%s height=%s hotspot=%d,%d hotspotX=%s hotspotY=%s\n", info.DataName, info.WidthName, info.HeightName, info.Hotspot.X, info.Hotspot.Y, info.HotspotXName, info.HotspotYName)
}

func ExampleDecodeInfoAndImage() {
	n := "cursor.xbm"
	f, err := os.Open(n)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	info, m, err := DecodeInfoAndImage(f)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Decoded XBM Dimensions: %dx%d\n", m.Bounds().Max.X, m.Bounds().Max.Y)
	fmt.Printf("Metadata: data=%s width=%s height=%s hotspot=%d,%d hotspotX=%s hotspotY=%s\n", info.DataName, info.WidthName, info.HeightName, info.Hotspot.X, info.Hotspot.Y, info.HotspotXName, info.HotspotYName)
}
