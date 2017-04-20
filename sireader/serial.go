package sireader


type Serial struct {
	port string
	baudrate int
	timeout int
}

func NewSerial(port string) *Serial {
	return &Serial{port: port}
}
