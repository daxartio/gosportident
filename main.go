package main

import (
	sr "./sireader"
)

func main() {
	r, _ := sr.NewReader("COM3")
	r.Beep()
}
