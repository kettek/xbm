package xbm

import (
	"image"
	"os"
	"testing"
)

func TestWrite(t *testing.T) {
	encoder := Encoder{}
	encoder.DataName = "datums"
	encoder.Hotspot.X = 2
	f, err := os.Create("tests/_write.xbm")
	if err != nil {
		t.Error(err)
	}
	encoder.Encode(f, image.NewRGBA(image.Rect(0, 0, 4, 4)))
}
