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
	c := &serial.Config{Name: port, Baud: 34800}
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
	return ProtoConfig{}
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

func toBytes(data interface{}) ([]byte, error) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, data)
	if err == nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func crc(b []byte) []byte {
	toChars := func (s []byte)([]byte, error) {
		if len(s) == 0 {
			return nil, nil
		}
		if len(s)%2 == 0 {
			s = append(s, []byte{0x00, 0x00}...)
		} else {
			s = append(s, 0x00)
		}
		// TODO cделать генератор

		return []byte{}, nil
	}
	//crc := toInt(b[:2])
	toChars(b[2:])

	return []byte{}
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

func (r *Reader) sendCommand(command []byte) (int, error){
	cmd := append(STX, command...)
	cmd = append(cmd, ETX...)

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
