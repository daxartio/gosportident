package sireader


type ProtoConfig struct {
	ExtProto byte
	AutoSend byte
	HandShake byte
	PWAccess byte
	PunchRead byte
}


type Reader struct {
	serial *Serial
	debug bool
	logfile string
	protoConfig ProtoConfig
}

func (r *Reader) SetExtendedProtocol(extended bool)

func (r *Reader) SetAutosend(autosend bool)

func (r *Reader) SetOperatingMode(mode string)

func (r *Reader) SetStationCode(code int)

func (r *Reader) GetTime() DateTime

func (r *Reader) SetTime()

func (r *Reader) Beep()

func (r *Reader) PowerOff()

func (r *Reader) Disconnect()

func (r *Reader) Reconnect()

func (r *Reader) connectReader()

func (r *Reader) updateProtoConfig() ProtoConfig

func (r *Reader) SetProtoConfig(config ProtoConfig)

func toInt(s string) int

func toString(i int, len int) string

func crc(s string) byte

func crcCheck(s string, crc string) bool

func decodeCardnr(number int) int

func decodeTime()

func appendPunch()

func decodeCardData()

func (r *Reader) sendCommand()

func (r *Reader) readCommand()


type ReaderReadout struct {
	Reader
	card string
	CardType string
}

func (rr *ReaderReadout) PollCard()

func (rr *ReaderReadout) ReadCard(reftime int)

func (rr *ReaderReadout) AckCard()

func (rr *ReaderReadout) readCommand(timeout int) 


type ReaderControl struct {
	Reader
	nextOffset string
}

func (rc *ReaderControl) PolPunch()

func (rc *ReaderControl) readPunch(offset int)
