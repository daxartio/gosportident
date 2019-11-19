package sireader

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

func crcCheck(sample, expectedCrc []byte) bool {
	return toInt(crc(sample)) == toInt(expectedCrc)
}
