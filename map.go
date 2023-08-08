package packed

// Map represents a map. The encoded value is a sequence of key-value pairs, where each key and value
// is encoded using the provided key and value type encoders.
func Map[K comparable, V any](keys Type[K], values Type[V]) Type[map[K]V] {
	return func(ptr *map[K]V) Packed {
		return NewEncodeDecode(
			// Encode
			func() ([]byte, error) {
				buf := []byte{byte(len(*ptr))}
				for k, v := range *ptr {
					// Encode the key and add to the buffer
					keyEncoded, err := keys(&k).Encode()
					if err != nil {
						return nil, err
					}
					buf = append(buf, keyEncoded...)

					// Encode the value and add to the buffer
					valueEncoded, err := values(&v).Encode()
					if err != nil {
						return nil, err
					}
					buf = append(buf, valueEncoded...)
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

				*ptr = make(map[K]V)
				for i := 0; i < length; i++ {
					// Decode the key
					var key K
					n, err := keys(&key).Decode(buf[cursor:])
					if err != nil {
						return 0, err
					}
					cursor += n

					// Decode the value
					var value V
					n, err = values(&value).Decode(buf[cursor:])
					if err != nil {
						return 0, err
					}
					cursor += n

					// Add the key and value to the map
					(*ptr)[key] = value
				}
				return cursor, nil
			},
		)
	}
}
