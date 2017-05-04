package sireader

import (
	"testing"
	"fmt"
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

func TestToInt(t *testing.T) {
	println("TestToInt")
	variable := []testBytesToInt{
		{[]byte("\xee\xee"), 61166},
		{TIME_RESET, 61166},
		{[]byte{'\x00'}, 0},
		{BC_SI5_DET, 70},
		{BC_SI6_WRITEPAGE, 98},
	}
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