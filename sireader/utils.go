package sireader

import (
	"bytes"
	"encoding/binary"
)

//BytesMerge is func merging slices bytes
func BytesMerge(args ...[]byte) []byte {
	count := 0
	for _, arr := range args {
		count += len(arr)
	}
	bs := make([]byte, count)
	i := 0
	for _, arr := range args {
		for _, el := range arr {
			bs[i] = el
			i++
		}
	}
	return bs
}

func toInt(s []byte) int {
	value := 0
	for _, c := range s {
		value = value*256 + int(c)
	}
	return value
}

func toBytes(data int) []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, uint32(data))
	b := bytes.Trim(buf.Bytes(), "\x00")
	bLength := len(b)
	result := make([]byte, bLength)
	for i := 0; i < bLength; i++ {
		result[i] = b[bLength-1-i]
	}
	return result
}
