package packed

// Bool represents a boolean value. The encoded value is 0 if false, and 1 if true.
func Bool(ptr *bool) Packed {
	return NewEncodeDecode(
		// Encode
		func() ([]byte, error) {
			buf := make([]byte, 1)
			if *ptr {
				buf[0] = 1
			}
			return buf, nil
		},
		// Decode
		func(buf []byte) (int, error) {
			if len(buf) < 1 {
				return 0, ErrorInvalidLength
			}
			*ptr = buf[0] != 0
			return 1, nil
		},
	)
}
