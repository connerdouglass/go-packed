package packed

import (
	"encoding/binary"
	"math"
)

// Float32 represents a float32 value. The encoded value is the IEEE 754 binary representation of the float32.
func Float32(ptr *float32) Packed {
	return fixedSize[float32](
		ptr, 4,
		func(b []byte, f float32) {
			binary.LittleEndian.PutUint32(b, math.Float32bits(f))
		},
		func(b []byte) float32 {
			return math.Float32frombits(binary.LittleEndian.Uint32(b))
		},
	)
}

// Float64 represents a float64 value. The encoded value is the IEEE 754 binary representation of the float64.
func Float64(ptr *float64) Packed {
	return fixedSize[float64](
		ptr, 8,
		func(b []byte, f float64) {
			binary.LittleEndian.PutUint64(b, math.Float64bits(f))
		},
		func(b []byte) float64 {
			return math.Float64frombits(binary.LittleEndian.Uint64(b))
		},
	)
}
