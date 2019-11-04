package main

import (
	"fmt"

	sr "./sireader"
)

func main() {
	fmt.Println("Hello")
	r, _ := sr.NewReader("COM1")
	r.Beep()
	fmt.Println(r.GetTime())
	fmt.Println("END")
}
