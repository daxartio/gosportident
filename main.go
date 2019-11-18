package main

import (
	"fmt"

	sr "./sireader"
	// "log"

	// "github.com/tarm/serial"
)

func main() {
	fmt.Println("Hello")
	r, _ := sr.NewReader("/dev/ttyUSB0")
	r.Beep()
	fmt.Println("END")
	
	// beepTwiceCmd := []byte{0x02, 0xF9, 0x01, 0x02, 0x14, 0x0A, 0x03}
	// si10Cmd := []byte{0x02, 0xEF, 0x01, 0x01, 0xEA, 0x09, 0x03}

	// c := &serial.Config{Name: "/dev/ttyUSB0", Baud: 38400}
	// s, err := serial.OpenPort(c)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	
	// n, err := s.Write(beepTwiceCmd)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	
	// buf := make([]byte, 256)
	// n, err = s.Read(buf)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Printf("%q", buf[:n])
	
	// buf = make([]byte, 256)
	// n, err = s.Read(buf)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Printf("%q", buf[:n])
	
	// n, err = s.Write(si10Cmd)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Printf("%q", buf[:n])
	
	// buf = make([]byte, 256)
	// n, err = s.Read(buf)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Printf("%q", buf[:n])
	
	// buf = make([]byte, 256)
	// n, err = s.Read(buf)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Printf("%q", buf[:n])
	
	// n, err = s.Write(beepTwiceCmd)
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
