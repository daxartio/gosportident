package sireader

import (
	"testing"
)

func TestNewReader(t *testing.T) {
	reader, err := NewReader("COM45")
	if err != nil {
		t.Fatal(err)
	}
	reader.debug = true
}

func TestToInt(t *testing.T)

func TestToString(t *testing.T)

func TestCrc(t *testing.T)

func TestCrcCheck(t *testing.T)

func TestDecodeTime(t *testing.T)