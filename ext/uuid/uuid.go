package uuid

import (
	"github.com/connerdouglass/go-packed"
	"github.com/google/uuid"
)

// UUID represents a UUID value.
func UUID(ptr *uuid.UUID) packed.Packed {
	return packed.NewEncodeDecode(
		// Encode
		func() ([]byte, error) {
			buf := make([]byte, 16)
			copy(buf, ptr[:])
			return buf, nil
		},
		// Decode
		func(buf []byte) (int, error) {
			if len(buf) < 16 {
				return 0, packed.ErrorInvalidLength
			}
			copy(ptr[:], buf)
			return 16, nil
		},
	)
}
