package xbm

import (
	"bufio"
	"fmt"
	"image"
	"io"
)

// decoder is the type used to decode XBM data.
type decoder struct {
	scanner            *bufio.Scanner
	pixelData          []byte
	width, height      int
	hotspotX, hotspotY int
}

func (d *decoder) parseHeader() error {
	for d.scanner.Scan() {
		// TODO
		fmt.Println(d.scanner.Text())
	}
	if err := d.scanner.Err(); err != nil {
		return err
	}
	return nil
}

// parsePixels parses the XBM pixel data.
func (d *decoder) parsePixels() error {
	for d.scanner.Scan() {
		// TODO
		fmt.Println(d.scanner.Text())
	}
	if err := d.scanner.Err(); err != nil {
		return err
	}
	return nil
}

// Decode reads an XBM file from r and returns the image.
func Decode(r io.Reader) (image.Image, error) {
	d := &decoder{
		scanner: bufio.NewScanner(r),
	}
	if err := d.parseHeader(); err != nil {
		return nil, err
	}
	if err := d.parsePixels(); err != nil {
		return nil, err
	}
	// TODO
	return nil, nil
}

// DecodeConfig returns the dimensions of an XBM image without
// decoding the entire image.
func DecodeConfig(r io.Reader) (image.Config, error) {
	d := &decoder{
		scanner: bufio.NewScanner(r),
	}
	if err := d.parseHeader(); err != nil {
		return image.Config{}, err
	}
	// TODO
	return image.Config{
		BitColorModel, d.width, d.height,
	}, nil
}

func init() {
	image.RegisterFormat("xbm", "/*", Decode, DecodeConfig)
	image.RegisterFormat("xbm", "#define", Decode, DecodeConfig)
}
