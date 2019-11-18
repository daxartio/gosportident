package sireader

import (
	"fmt"
	"log"
	"time"

	"github.com/tarm/serial"
)

//ProtoConfig is configuration struct
type ProtoConfig struct {
	ExtProto  byte
	AutoSend  byte
	HandShake byte
	PWAccess  byte
	PunchRead byte
}

//Reader is main struct
type Reader struct {
	port        *serial.Port
	protoConfig ProtoConfig
	debug       bool
	logfile     string
}

//NewReader is constructor of Reader
func NewReader(port string) (*Reader, error) {
	c := &serial.Config{Name: port, Baud: 38400, ReadTimeout: time.Second * 5}
	s, err := serial.OpenPort(c)
	if err != nil {
		return nil, err
	}
	r := new(Reader)
	r.port = s
	// r.ConnectReader()
	return r, nil
}

//func (r *Reader) SetExtendedProtocol(extended bool) {}

//func (r *Reader) SetAutoSend(autoSend bool) {}

//func (r *Reader) SetOperatingMode(mode string) {}

//func (r *Reader) SetStationCode(code int) {}

//GetTime is returning time object
func (r *Reader) GetTime() time.Time {
	bintTime, e := r.sendCommand([]byte{CGetTime}, []byte{})
	fmt.Println(bintTime, e)
	return time.Time{}
}

//func (r *Reader) SetTime(t *time.Time) {}

//Beep cmd for si station
func (r *Reader) Beep() {
	r.sendCommand([]byte{CBeep}, toBytes(1))
}

//func (r *Reader) PowerOff() {}

//func (r *Reader) Disconnect() {}

//func (r *Reader) Reconnect() {}

//ConnectReader connect reading
func (r *Reader) ConnectReader() {
	r.port.Flush()
	r.sendCommand([]byte{CSetMs}, []byte{PMsDirect})
}

func (r *Reader) updateProtoConfig() ProtoConfig {
	//ret, _ := r.sendCommand(Bytes(C_GET_SYS_VAL),Bytes(O_PROTO, 0x01))
	//configByte := uint(toInt(Bytes()))
	protoConfig := ProtoConfig{}
	//protoConfig.ExtProto = configByte & (1 << 0) != 0
	//protoConfig.AutoSend = configByte & (1 << 1) != 0
	//protoConfig.HandShake = configByte & (1 << 2) != 0
	//protoConfig.PWAccess = configByte & (1 << 4) != 0
	//protoConfig.PunchRead = configByte & (1 << 7) != 0
	return protoConfig
}

//func (r *Reader) SetProtoConfig(config ProtoConfig) {}

func crc(b []byte) []byte {
	toChars := func(s []byte) [][]byte {
		if len(s) == 0 {
			return nil
		}
		if len(s)%2 == 0 {
			s = append(s, []byte{0x00, 0x00}...)
		} else {
			s = append(s, 0x00)
		}
		var result [][]byte
		for i := 0; i < len(s); i++ {
			result = append(result, s[i:i+2])
			i++
		}
		return result
	}

	if len(b) < 1 {
		return []byte{0x00, 0x00}
	}
	if len(b) == 2 {
		return b
	}

	crc := uint16(toInt(b[:2]))
	ch := toChars(b[2:])
	for _, c := range ch {
		val := uint16(toInt(c))
		for j := 0; j < 16; j++ {
			if (crc & CrcBitf) != 0 {
				crc <<= 1

				if (val & CrcBitf) != 0 {
					crc++
				}

				crc ^= CrcPolynom
			} else {
				crc <<= 1

				if (val & CrcBitf) != 0 {
					crc++
				}
			}
			val <<= 1
		}
	}

	crc &= 0xFFFF
	return BytesMerge(toBytes(int(crc>>8)), toBytes(int(crc&0xFF)))
}

func crcCheck(s string, crc string) bool {
	return false
}

func decodeCardNr(number int) int {
	return 0
}

//func decodeTime() {}

//func appendPunch() {}

//func decodeCardData() {}

func (r *Reader) sendCommand(command, parameters []byte) (int, error) {
	cmd := BytesMerge(command, toBytes(len(parameters)), parameters)
	cmd = BytesMerge([]byte{STX}, cmd, crc(cmd), []byte{ETX})

	n, e := r.port.Write(cmd)
	r.readCommand()
	return n, e
}

func (r *Reader) readCommand() {
	buf := make([]byte, 128)
	n, err := r.port.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%q", buf[:n])
}

//ReaderReadout is struct for readout
type ReaderReadout struct {
	Reader
	card     string
	CardType string
}

//func (rr *ReaderReadout) PollCard() {}

//func (rr *ReaderReadout) ReadCard(refTime int) {}

//func (rr *ReaderReadout) AckCard() {}

//func (rr *ReaderReadout) readCommand(timeout int) {}

//ReaderControl is struct for control
type ReaderControl struct {
	Reader
	nextOffset string
}

//func (rc *ReaderControl) PolPunch() {}

//func (rc *ReaderControl) readPunch(offset int) {}
