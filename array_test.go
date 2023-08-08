package packed_test

import (
	"testing"

	"github.com/connerdouglass/go-packed"
	"github.com/stretchr/testify/require"
)

func TestPackedBytes(t *testing.T) {
	t.Run("arrays", func(t *testing.T) {
		type testType struct {
			Bytes1 [4]byte
			Bytes2 [7]byte
		}
		val := testType{
			Bytes1: [4]byte{1, 3, 7, 9},
			Bytes2: [7]byte{2, 4, 6, 8, 10, 12, 14},
		}
		p := packed.Sequence(
			packed.Array(packed.Byte)(val.Bytes1[:]),
			packed.Array(packed.Byte)(val.Bytes2[:]),
		)
		encoded, err := p.Encode()
		require.NoError(t, err, "encoding arrays test")
		require.NotNil(t, encoded, "encoded arrays test")

		val.Bytes1 = [4]byte{}
		val.Bytes2 = [7]byte{}

		n, err := p.Decode(encoded)
		require.NoError(t, err, "decoding arrays test")
		require.Equal(t, len(encoded), n, "decoded arrays test")

		require.Equal(t, [4]byte{1, 3, 7, 9}, val.Bytes1, "decoded arrays test")
		require.Equal(t, [7]byte{2, 4, 6, 8, 10, 12, 14}, val.Bytes2, "decoded arrays test")
	})
	t.Run("slices", func(t *testing.T) {
		type testType struct {
			Bytes1   []byte
			Bytes2   []byte
			Strings1 []string
			NilSlice []uint32
		}
		val := testType{
			Bytes1:   []byte{1, 3, 7, 9},
			Bytes2:   []byte{2, 4, 6, 8, 10, 12, 14},
			Strings1: []string{"hello", "world"},
			NilSlice: nil,
		}
		p := packed.Sequence(
			packed.Slice(packed.Byte)(&val.Bytes1),
			packed.Slice(packed.Byte)(&val.Bytes2),
			packed.Slice(packed.String)(&val.Strings1),
			packed.Slice(packed.Uint32)(&val.NilSlice),
		)
		encoded, err := p.Encode()
		require.NoError(t, err, "encoding slices test")
		require.NotNil(t, encoded, "encoded slices test")

		val.Bytes1 = nil
		val.Bytes2 = nil
		val.Strings1 = nil
		val.NilSlice = []uint32{1, 2}

		n, err := p.Decode(encoded)
		require.NoError(t, err, "decoding slices test")
		require.Equal(t, len(encoded), n, "decoded slices test")

		require.Equal(t, []byte{1, 3, 7, 9}, val.Bytes1, "decoded slices test")
		require.Equal(t, []byte{2, 4, 6, 8, 10, 12, 14}, val.Bytes2, "decoded slices test")
		require.Equal(t, []string{"hello", "world"}, val.Strings1, "decoded slices test")
		require.Equal(t, []uint32{}, val.NilSlice, "decoded slices test")
	})
}
