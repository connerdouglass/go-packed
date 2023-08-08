package packed

import "encoding/binary"

// String represents a string value. The first two bytes of an encoded string is the length
// of the string. The remaining bytes encode the string.
func String(ptr *string) Packed {
	return NewEncodeDecode(
		// Encode
		func() ([]byte, error) {
			buf := make([]byte, 2+len([]byte(*ptr)))
			binary.LittleEndian.PutUint16(buf, uint16(len(*ptr)))
			copy(buf[2:], []byte(*ptr))
			return buf, nil
		},
		// Decode
		func(buf []byte) (int, error) {
			if len(buf) < 2 {
				return 0, ErrorInvalidLength
			}
			strLen := int(binary.LittleEndian.Uint16(buf))
			if len(buf) < 2+strLen {
				return 0, ErrorInvalidLength
			}
			*ptr = string(buf[2 : 2+strLen])
			return 2 + strLen, nil
		},
	)
}

// String256 represents a string value up to 256 characters in length. The first byte of an encoded string is the length
// of the string. The remaining bytes encode the string.
func String256(ptr *string) Packed {
	return NewEncodeDecode(
		// Encode
		func() ([]byte, error) {
			buf := make([]byte, 1+len([]byte(*ptr)))
			buf[0] = uint8(len(*ptr))
			copy(buf[1:], []byte(*ptr))
			return buf, nil
		},
		// Decode
		func(buf []byte) (int, error) {
			if len(buf) < 1 {
				return 0, ErrorInvalidLength
			}
			strLen := int(buf[0])
			if len(buf) < 1+strLen {
				return 0, ErrorInvalidLength
			}
			*ptr = string(buf[1 : 1+strLen])
			return 1 + strLen, nil
		},
	)
}
