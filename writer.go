package xbm

import (
	"fmt"
	"image"
	"io"
)

// Encoder configures additional XBM encoding data.
type Encoder struct {
	Info
}

type encoder struct {
	enc *Encoder
	w   io.Writer
	m   image.Image
}

// writeHeader writes the XBM header information.
func (enc *encoder) writeHeader() error {
	widthName := "xbm_width"
	heightName := "xbm_height"
	hotspotXName := "xbm_hotspot_x"
	hotspotYName := "xbm_hotspot_y"
	useHotspots := false
	width := enc.m.Bounds().Max.X - enc.m.Bounds().Min.X
	height := enc.m.Bounds().Max.Y - enc.m.Bounds().Min.Y

	if enc.enc.WidthName != "" {
		widthName = enc.enc.WidthName
	}
	if enc.enc.HeightName != "" {
		heightName = enc.enc.HeightName
	}
	if enc.enc.HotspotXName != "" {
		hotspotXName = enc.enc.HotspotXName
		useHotspots = true
	}
	if enc.enc.HotspotYName != "" {
		hotspotYName = enc.enc.HotspotYName
		useHotspots = true
	}

	// Write out width and height.
	enc.w.Write([]byte(fmt.Sprintf("#define %s %d\n", widthName, width)))
	enc.w.Write([]byte(fmt.Sprintf("#define %s %d\n", heightName, height)))

	// Write out hotspots if desired.
	if enc.enc.Hotspot.X > 0 || enc.enc.Hotspot.Y > 0 || useHotspots {
		enc.w.Write([]byte(fmt.Sprintf("#define %s %d\n", hotspotXName, enc.enc.Hotspot.X)))
		enc.w.Write([]byte(fmt.Sprintf("#define %s %d\n", hotspotYName, enc.enc.Hotspot.Y)))
	}
	return nil
}

// writePixels writes the XBM pixel data.
func (enc *encoder) writePixels() error {
	dataName := "xbm_bits"
	if enc.enc.DataName != "" {
		dataName = enc.enc.DataName
	}

	enc.w.Write([]byte(fmt.Sprintf("static unsigned char %s[] = {\n", dataName)))
	// ...
	enc.w.Write([]byte(fmt.Sprint("\n};")))
	// TODO
	return nil
}

// Encode writes the Info i and Image image to w in PNG format. Any Image may be
// encoded, but images that are not Bits are encoded lossily.
func Encode(w io.Writer, i Info, image image.Image) error {
	e := Encoder{
		i,
	}
	return e.Encode(w, image)
}

// Encode write the Image m to w in XBM format. Any Image may be
// encoded, but images that are not Bits are encoded lossily.
func (enc *Encoder) Encode(w io.Writer, m image.Image) error {
	e := encoder{
		enc: enc,
		w:   w,
		m:   m,
	}
	if err := e.writeHeader(); err != nil {
		return err
	}
	if err := e.writePixels(); err != nil {
		return err
	}
	return nil
}
