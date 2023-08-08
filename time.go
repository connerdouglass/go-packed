package packed

import (
	"encoding/binary"
	"time"
)

// Time represents a time value. The encoded value is the number of seconds since the Unix epoch.
func Time(ptr *time.Time) Packed {
	return NewEncodeDecode(
		// Encode
		func() ([]byte, error) {
			buf := make([]byte, 8)
			binary.LittleEndian.PutUint64(buf, uint64(ptr.Unix()))
			return buf, nil
		},
		// Decode
		func(buf []byte) (int, error) {
			if len(buf) < 8 {
				return 0, ErrorInvalidLength
			}
			*ptr = time.Unix(int64(binary.LittleEndian.Uint64(buf)), 0)
			return 8, nil
		},
	)
}
