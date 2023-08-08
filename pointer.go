package packed

// Pointer represents a pointer to a value. The first byte of an encoded pointer is
// 0 if the pointer is nil, and 1 otherwise. The remaining bytes encode the dereferenced value.
func Pointer[T any](child Type[T]) Type[*T] {
	return func(ptr **T) Packed {
		return NewEncodeDecode(
			// Encode
			func() ([]byte, error) {
				if *ptr == nil {
					return []byte{0}, nil
				}
				buf := []byte{1}
				innerEncoded, err := child(*ptr).Encode()
				if err != nil {
					return nil, err
				}
				buf = append(buf, innerEncoded...)
				return buf, nil
			},
			// Decode
			func(buf []byte) (int, error) {
				if len(buf) < 1 {
					return 0, ErrorInvalidLength
				}
				if buf[0] == 0 {
					*ptr = nil
					return 1, nil
				}
				*ptr = new(T)
				n, err := child(*ptr).Decode(buf[1:])
				if err != nil {
					return 1, err
				}
				return 1 + n, nil
			},
		)
	}
}
