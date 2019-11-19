package main

import (
	"../../sireader"
	"fmt"
	"log"
)

func main() {
	// r, err := sireader.NewReader("COM3")
	r, err := sireader.NewReader("/dev/ttyUSB0")
	if err != nil {
		log.Fatal(err)
	}
	r.Beep()
	fmt.Println(r.GetTime())
	r.Poll()
	err = r.Disconnect()
	if err != nil {
		log.Fatal(err)
	}
}
