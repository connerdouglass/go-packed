package packed

// Slice represents a slice of values. The first byte of an encoded slice is the length of the slice.
func Slice[T any](children Type[T]) Type[[]T] {
	return func(ptr *[]T) Packed {
		return NewEncodeDecode(
			// Encode
			func() ([]byte, error) {
				buf := []byte{
					byte(len(*ptr)),
				}
				for i := 0; i < len(*ptr); i++ {
					childEncoded, err := children(&(*ptr)[i]).Encode()
					if err != nil {
						return nil, err
					}
					buf = append(buf, childEncoded...)
				}
				return buf, nil
			},
			// Decode
			func(buf []byte) (int, error) {
				if len(buf) < 1 {
					return 0, ErrorInvalidLength
				}
				var cursor int

				length := int(buf[cursor])
				cursor++

				*ptr = make([]T, length)
				for i := 0; i < length; i++ {
					n, err := children(&(*ptr)[i]).Decode(buf[cursor:])
					if err != nil {
						return 0, err
					}
					cursor += n
				}
				return cursor, nil
			},
		)
	}
}

// Array represents an array of values. The encoded values are concatenated. Although this can technically be used
// with slices, you should only use it if the length is truly fixed, and known at compile time.
func Array[T any](children Type[T]) func([]T) Packed {
	return func(ptr []T) Packed {
		size := len(ptr)
		return NewEncodeDecode(
			// Encode
			func() ([]byte, error) {
				var buf []byte
				for i := 0; i < size; i++ {
					childEncoded, err := children(&ptr[i]).Encode()
					if err != nil {
						return nil, err
					}
					buf = append(buf, childEncoded...)
				}
				return buf, nil
			},
			// Decode
			func(buf []byte) (int, error) {
				var cursor int
				for i := 0; i < size; i++ {
					n, err := children(&ptr[i]).Decode(buf[cursor:])
					if err != nil {
						return 0, err
					}
					cursor += n
				}
				return cursor, nil
			},
		)
	}
}
