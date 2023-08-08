package packed

import "encoding/binary"

func fixedSize[T any](
	ptr *T,
	size int,
	encodeFunc func([]byte, T),
	decodeFunc func([]byte) T,
) Packed {
	return NewEncodeDecode(
		// Encode
		func() ([]byte, error) {
			buf := make([]byte, size)
			encodeFunc(buf, *ptr)
			return buf, nil
		},
		// Decode
		func(buf []byte) (int, error) {
			if len(buf) < size {
				return 0, ErrorInvalidLength
			}
			*ptr = decodeFunc(buf)
			return size, nil
		},
	)
}

func Byte(ptr *byte) Packed {
	return Uint8(ptr)
}

func Uint8(ptr *uint8) Packed {
	return fixedSize[uint8](
		ptr, 1,
		func(buf []byte, val uint8) {
			buf[0] = val
		},
		func(buf []byte) uint8 {
			return buf[0]
		},
	)
}

func Uint16(ptr *uint16) Packed {
	return fixedSize[uint16](ptr, 2, binary.LittleEndian.PutUint16, binary.LittleEndian.Uint16)
}

func Uint32(ptr *uint32) Packed {
	return fixedSize[uint32](ptr, 4, binary.LittleEndian.PutUint32, binary.LittleEndian.Uint32)
}

func Uint64(ptr *uint64) Packed {
	return fixedSize[uint64](ptr, 8, binary.LittleEndian.PutUint64, binary.LittleEndian.Uint64)
}

func Int8(ptr *int8) Packed {
	return fixedSize[int8](
		ptr, 1,
		func(buf []byte, val int8) {
			buf[0] = byte(val)
		},
		func(buf []byte) int8 {
			return int8(buf[0])
		},
	)
}

func Int16(ptr *int16) Packed {
	return fixedSize[int16](
		ptr, 2,
		func(buf []byte, val int16) {
			binary.LittleEndian.PutUint16(buf, uint16(val))
		},
		func(buf []byte) int16 {
			return int16(binary.LittleEndian.Uint16(buf))
		},
	)
}

func Int32(ptr *int32) Packed {
	return fixedSize[int32](
		ptr, 4,
		func(buf []byte, val int32) {
			binary.LittleEndian.PutUint32(buf, uint32(val))
		},
		func(buf []byte) int32 {
			return int32(binary.LittleEndian.Uint32(buf))
		},
	)
}

func Int64(ptr *int64) Packed {
	return fixedSize[int64](
		ptr, 8,
		func(buf []byte, val int64) {
			binary.LittleEndian.PutUint64(buf, uint64(val))
		},
		func(buf []byte) int64 {
			return int64(binary.LittleEndian.Uint64(buf))
		},
	)
}
