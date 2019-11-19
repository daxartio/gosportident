package main

import (
	"../../sireader"
	"fmt"
	"log"
)

func main() {
	r, err := sireader.NewReader("COM3")
	//r, _ := sireader.NewReader("/dev/ttyUSB0")
	if err != nil {
		log.Fatal(err)
	}
	r.Beep()
	fmt.Println(r.GetTime())
	err = r.Disconnect()
	if err != nil {
		log.Fatal(err)
	}
}
