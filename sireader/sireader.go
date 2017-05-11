package sireader

import (
	"time"
	"github.com/tarm/serial"
	"encoding/binary"
	"bytes"
)

type ProtoConfig struct {
	ExtProto  byte
	AutoSend  byte
	HandShake byte
	PWAccess  byte
	PunchRead byte
}

type Reader struct {
	port        *serial.Port
	protoConfig ProtoConfig
	debug       bool
	logfile     string
}

func NewReader(port string) (*Reader, error) {
	c := &serial.Config{Name: port, Baud: 38400, ReadTimeout: time.Second * 5}
	s, err := serial.OpenPort(c)
	if err != nil {
		return nil, err
	}

	return &Reader{port: s}, nil
}

func (r *Reader) SetExtendedProtocol(extended bool) {

}

func (r *Reader) SetAutoSend(autoSend bool) {

}

func (r *Reader) SetOperatingMode(mode string) {

}

func (r *Reader) SetStationCode(code int) {

}

func (r *Reader) GetTime() time.Time {
	return time.Time{}
}

func (r *Reader) SetTime(t *time.Time) {

}

func (r *Reader) Beep() {

}

func (r *Reader) PowerOff() {

}

func (r *Reader) Disconnect() {

}

func (r *Reader) Reconnect() {

}

func (r *Reader) connectReader() {

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

func (r *Reader) SetProtoConfig(config ProtoConfig) {

}

func toInt(s []byte) int {
	value := uint(0)
	for offset, c := range s {
		value += uint(c) << (uint(offset) * uint(8))
	}
	return int(value)
}

func toBytes(data int) []byte {
	buf := new(bytes.Buffer)
	num := uint64(data)
	binary.Write(buf, binary.LittleEndian, num)
	return buf.Bytes()
}

func crc(b []byte) []byte {
	toChars := func(s []byte) ([][]byte, error) {
		if len(s) == 0 {
			return nil, nil
		}
		if len(s)%2 == 0 {
			s = append(s, []byte{0x00, 0x00}...)
		} else {
			s = append(s, 0x00)
		}
		// TODO cделать генератор
		result := [][]byte{}
		for i := 0; i < len(s); i++ {
			result = append(result, s[i:i+2])
			i++
		}
		return result, nil
	}

	if len(b) < 1 {
		return Bytes(0x00, 0x00)
	}
	if len(b) == 2 {
		return b
	}

	crc := uint(toInt(b[:2]))
	ch, _ := toChars(b[2:])
	for _, c := range ch {
		val := uint(toInt(c))
		for j := 0; j < 16; j++ {
			if (crc & CRC_BITF) != 0 {
				crc <<= 1

				if (val & CRC_BITF) != 0 {
					crc += 1
				}

				crc ^= CRC_POLYNOM
			} else {
				crc <<= 1

				if (val & CRC_BITF) != 0 {
					crc += 1
				}
			}
			val <<= 1
		}
	}

	crc &= 0xFFFF
	return BytesMerge(bytes.Trim(toBytes(int(crc>>8)), "\x00"),
		bytes.Trim(toBytes(int(crc&0xFF)), "\x00"))
}

func crcCheck(s string, crc string) bool {
	return false
}

func decodeCardnr(number int) int {
	return 0
}

func decodeTime() {

}

func appendPunch() {

}

func decodeCardData() {

}

func (r *Reader) sendCommand(command, parameters []byte) (int, error) {
	cmd := BytesMerge(command, toBytes(len(parameters)), parameters)
	cmd = BytesMerge(Bytes(STX), cmd, crc(cmd), Bytes(ETX))

	return r.port.Write(cmd)
}

func (r *Reader) readCommand() {

}

type ReaderReadout struct {
	Reader
	card     string
	CardType string
}

func (rr *ReaderReadout) PollCard() {

}

func (rr *ReaderReadout) ReadCard(refTime int) {

}

func (rr *ReaderReadout) AckCard() {

}

func (rr *ReaderReadout) readCommand(timeout int) {

}

type ReaderControl struct {
	Reader
	nextOffset string
}

func (rc *ReaderControl) PolPunch() {

}

func (rc *ReaderControl) readPunch(offset int) {

}
