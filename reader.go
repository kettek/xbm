package xbm

import (
	"bufio"
	"encoding/hex"
	"image"
	"io"
	"strconv"
	"strings"
)

// FormatError represents an invalid format error for the XBM format.
type FormatError string

// Error returns the string of the error.
func (e FormatError) Error() string { return "xbm: invalid format: " + string(e) }

// State machine constants
const (
	seenNothing = iota
	seenDefineWidth
	seenDefineHeight
	seenDefineHotspotX
	seenDefineHotspotY
	seenStatic
)

// decoder is the type used to decode XBM data.
type decoder struct {
	scanner       *bufio.Scanner
	image         *Bits
	info          Info
	width, height int
	lastSeen      int
}

func (d *decoder) parseHeader() error {
	for d.scanner.Scan() {
		token := d.scanner.Text()
		if strings.HasPrefix(token, "//") {
			// consume
			continue
		} else if strings.HasPrefix(token, "#define") && d.lastSeen < seenStatic {
			words := strings.Split(token, " ")
			if len(words) < 3 {
				return FormatError("too few words in #define")
			}
			value, err := strconv.Atoi(words[2])
			if err != nil {
				return err
			}
			if d.lastSeen == seenNothing {
				d.width = value
				d.info.WidthName = words[1]
				d.lastSeen = seenDefineWidth
			} else if d.lastSeen == seenDefineWidth {
				d.height = value
				d.info.HeightName = words[1]
				d.lastSeen = seenDefineHeight
			} else if d.lastSeen == seenDefineHeight {
				d.info.Hotspot.X = value
				d.info.HotspotXName = words[1]
				d.lastSeen = seenDefineHotspotX
			} else if d.lastSeen == seenDefineHotspotX {
				d.info.Hotspot.Y = value
				d.info.HotspotYName = words[1]
				d.lastSeen = seenDefineHotspotY
			}
		} else if strings.HasPrefix(token, "static") && d.lastSeen >= seenDefineHeight {
			words := strings.Split(token, " ")
			for _, word := range words {
				if strings.HasSuffix(word, "[]") {
					d.info.DataName = word[:len(word)-2]
					break
				}
			}
			d.lastSeen = seenStatic
			break
		} else {
			return FormatError("invalid XBM data")
		}
	}
	if err := d.scanner.Err(); err != nil {
		return err
	}
	return nil
}

// parsePixels parses the XBM pixel data.
func (d *decoder) parsePixels() error {
	d.image = NewBits(image.Rect(
		0, 0, d.width, d.height,
	))
	var x, y int
	for d.scanner.Scan() {
		token := d.scanner.Text()
		words := strings.Split(strings.Trim(token, "\t {};"), ",")
		for _, word := range words {
			var byte byte
			if word == "" {
				continue
			}
			if strings.HasPrefix(word, "0x") {
				bytes, err := hex.DecodeString(word[2:])
				if err != nil {
					return err
				}
				if len(bytes) != 1 {
					return FormatError("Invalid byte data")
				}
				byte = bytes[0]
			} else {
				return FormatError("Non-hex byte provided")
			}
			for i := 0; i < 8; i++ {
				if x > d.image.Bounds().Max.X {
					// Ignore the remainder.
					break
				}
				bit := (byte&(1<<i) > 0)
				d.image.SetBit(x, y, bit)
				x++
			}
			if x >= d.image.Bounds().Max.X {
				x = 0
				y++
			}
		}
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
	return d.image, nil
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
	return image.Config{
		ColorModel: BitColorModel,
		Width:      d.width,
		Height:     d.height,
	}, nil
}

// DecodeInfo returns the Info of an XBM image without
// decoding the entire image.
func DecodeInfo(r io.Reader) (Info, error) {
	d := &decoder{
		scanner: bufio.NewScanner(r),
	}
	if err := d.parseHeader(); err != nil {
		return Info{}, err
	}
	return d.info, nil
}

// DecodeInfoAndImage returns the Info and the Image of
// and XBM image.
func DecodeInfoAndImage(r io.Reader) (Info, image.Image, error) {
	d := &decoder{
		scanner: bufio.NewScanner(r),
	}
	if err := d.parseHeader(); err != nil {
		return Info{}, nil, err
	}
	if err := d.parsePixels(); err != nil {
		return d.info, nil, err
	}
	return d.info, d.image, nil
}

func init() {
	image.RegisterFormat("xbm", "/*", Decode, DecodeConfig)
	image.RegisterFormat("xbm", "#define", Decode, DecodeConfig)
}
