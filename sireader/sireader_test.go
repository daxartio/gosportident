package sireader

import (
	"testing"
	"fmt"
	"bytes"
	"log"
)

func TestNewReader(t *testing.T) {
	reader, err := NewReader("COM3")
	if err != nil {
		t.Fatal(err)
	}

	// FI
	buf := make([]byte, 128)
	n, err := reader.port.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(toInt(buf[:n]))

	buf = make([]byte, 128)
	n, err = reader.port.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(toInt(buf[:n]))

	// FO
	buf = make([]byte, 128)
	n, err = reader.port.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(toInt(buf[:n]))

	buf = make([]byte, 128)
	n, err = reader.port.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(toInt(buf[:n]))
}

type testBytesToInt struct {
	values []byte
	result int
}
var variable = []testBytesToInt{
		{[]byte("\xee\xee"), 61166},
		{TIME_RESET, 61166},
		{[]byte{0x00}, 0},
		{Bytes(BC_SI5_DET), 70},
		{Bytes(BC_SI6_WRITEPAGE), 98},
	}

func TestToInt(t *testing.T) {
	println("TestToInt")

	for _, v := range variable {
		res := toInt(v.values)
		if res != v.result {
			t.Error(
				"For", v.values,
				"expected", v.result,
				"got", res,
			)
		} else {
			fmt.Println(v.values, " Ok")
		}
	}
}

func TestToBytes(t *testing.T) {
	println("TestToBytes")
	for _, v := range variable {
		res:= toBytes(v.result)
		if bytes.Equal(res, v.values) {
			t.Error(
				"For", v.result,
				"expected", v.values,
				"got", res,
			)
		} else {
			println(v.result, res, " Ok")
		}
	}
}

func TestToString(t *testing.T) {
	println("TestToString")
	fmt.Println(toBytes(65535))
}

func TestCrc(t *testing.T) {

}

func TestCrcCheck(t *testing.T) {

}

func TestDecodeTime(t *testing.T) {

}