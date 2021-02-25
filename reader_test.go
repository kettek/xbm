package xbm

import (
	"os"
	"testing"
)

func TestReadAndPrint(t *testing.T) {
	_, err := ReadXBM("tests/goblin.xbm")
	if err != nil {
		t.Error(err)
	}
}

func ReadXBM(path string) (XBM, error) {
	f, err := os.Open(path)
	if err != nil {
		return XBM{}, err
	}
	defer f.Close()

	_, err = Decode(f)
	if err != nil {
		return XBM{}, err
	}

	return XBM{}, err
}
