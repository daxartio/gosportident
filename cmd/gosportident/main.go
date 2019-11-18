package main

import (
	"../../sireader"
)

func main() {
	r, _ := sireader.NewReader("COM3")
	//r, _ := sireader.NewReader("/dev/ttyUSB0")
	r.Beep()
}
