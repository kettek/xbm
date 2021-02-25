package main

import (
	"fmt"
	"image"
	"log"
	"os"

	"github.com/kettek/xbm"
)

func main() {
	if len(os.Args) < 2 {
		log.Println("XBM file argument required")
		return
	}
	info, m, err := decodeXBM(os.Args[1])
	if err != nil {
		log.Panic(err)
	}
	b := m.Bounds()
	fmt.Printf("Decoded XBM Dimensions: %dx%d->%dx%d\n", b.Min.X, b.Min.Y, b.Max.X, b.Max.Y)
	fmt.Printf("Metadata: data=%s width=%s height=%s hotspot=%d,%d hotspotX=%s hotspotY=%s\n", info.DataName, info.WidthName, info.HeightName, info.Hotspot.X, info.Hotspot.Y, info.HotspotXName, info.HotspotYName)
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			v, ok := m.At(x, y).(xbm.BitColor)
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

func decodeXBM(path string) (xbm.Info, image.Image, error) {
	f, err := os.Open(path)
	if err != nil {
		return xbm.Info{}, nil, err
	}
	defer f.Close()

	info, img, err := xbm.DecodeInfoAndImage(f)
	if err != nil {
		return info, nil, err
	}

	return info, img, err
}
