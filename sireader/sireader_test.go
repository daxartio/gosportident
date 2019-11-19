package sireader

import (
	"reflect"
	"testing"
)

func TestBytesToInt(t *testing.T) {
	sample := []byte{0x14, 0x0A}
	// sample := []byte("\x14\x0a")
	expectedInt := 5130

	sampleInt := toInt(sample)
	if sampleInt != expectedInt {
		t.Fatal("Converted bytes is not equal expected int", toInt(sample), expectedInt)
	}
}

func TestIntToBytes(t *testing.T) {
	sample := 5130
	expectedBytes := []byte{0x14, 0x0A}

	sampleBytes := toBytes(sample)
	if !reflect.DeepEqual(sampleBytes, expectedBytes) {
		t.Fatal("Converted int is not equal expected bytes", sampleBytes, expectedBytes)
	}
}

func TestCrcBeep(t *testing.T) {
	sample := []byte{CBeep, 0x01, 0x02}
	expectedCRC := []byte{0x14, 0x0A}

	sampleCRC := crc(sample)
	if !reflect.DeepEqual(sampleCRC, expectedCRC) {
		t.Fatal("Crc is not equal expected bytes", toInt(sampleCRC), toInt(expectedCRC))
	}
}

func TestReturnedCrcBeep(t *testing.T) {
	sample := []byte("\xF9\x03\x00\x04\x01")
	expectedCRC := []byte{0xF5, 0x0F}

	sampleCRC := crc(sample)
	if toInt(sampleCRC) != toInt(expectedCRC) {
		t.Fatal("Crc is not equal expected bytes", toInt(sampleCRC), toInt(expectedCRC))
	}
}

func TestSomeCrc(t *testing.T) {
	sample := []byte("\xF7\x00")
	expectedCRC := []byte{0xF7, 0x00}

	sampleCRC := crc(sample)
	if toInt(sampleCRC) != toInt(expectedCRC) {
		t.Fatal("Crc is not equal expected bytes", toInt(sampleCRC), toInt(expectedCRC))
	}
}

func TestCrc(t *testing.T) {
	sample := []byte{
		0x53, 0x00, 0x05, 0x01,
		0x0F, 0xB5, 0x00, 0x00,
		0x1E, 0x08,
	}
	expectedCRC := 0x2C12

	sampleCRC := crc(sample)
	if toInt(sampleCRC) != int(expectedCRC) {
		t.Fatal("Crc is not equal expected bytes", toInt(sampleCRC), int(expectedCRC))
	}
}

func TestBytesMerge(t *testing.T) {
	sample := [][]byte{[]byte{CBeep, 0x01, 0x02}, []byte{0x03}, []byte{}}
	expected := []byte{CBeep, 0x01, 0x02, 0x03}
	if !reflect.DeepEqual(BytesMerge(sample...), expected) {
		t.Fatal("Merged bytes is not equal expected bytes")
	}
}
