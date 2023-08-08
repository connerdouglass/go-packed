package packed_test

import (
	"testing"

	"github.com/connerdouglass/go-packed"
	"github.com/stretchr/testify/require"
)

func TestPackedInt(t *testing.T) {
	t.Run("unsigned integers", func(t *testing.T) {
		type testType struct {
			Uint16 uint16
			Uint64 uint64
			Uint32 uint32
		}
		val := testType{
			Uint16: 123,
			Uint64: 456,
			Uint32: 789,
		}
		p := packed.Sequence(
			packed.Uint16(&val.Uint16),
			packed.Uint32(&val.Uint32),
			packed.Uint64(&val.Uint64),
		)
		encoded, err := p.Encode()
		require.NoError(t, err, "encoding uint test")
		require.NotNil(t, encoded, "encoded uint test")

		val.Uint16 = 0
		val.Uint64 = 0
		val.Uint32 = 0

		n, err := p.Decode(encoded)
		require.NoError(t, err, "decoding uint test")
		require.Equal(t, len(encoded), n, "decoded uint test")

		require.Equal(t, uint16(123), val.Uint16, "decoded uint test")
		require.Equal(t, uint64(456), val.Uint64, "decoded uint test")
		require.Equal(t, uint32(789), val.Uint32, "decoded uint test")
	})
	t.Run("signed integers", func(t *testing.T) {
		type testType struct {
			Int16 int16
			Int64 int64
			Int32 int32
		}
		val := testType{
			Int16: -123,
			Int64: 456,
			Int32: -789,
		}
		p := packed.Sequence(
			packed.Int16(&val.Int16),
			packed.Int32(&val.Int32),
			packed.Int64(&val.Int64),
		)
		encoded, err := p.Encode()
		require.NoError(t, err, "encoding int test")
		require.NotNil(t, encoded, "encoded int test")

		val.Int16 = 0
		val.Int64 = 0
		val.Int32 = 0

		n, err := p.Decode(encoded)
		require.NoError(t, err, "decoding int test")
		require.Equal(t, len(encoded), n, "decoded int test")

		require.Equal(t, int16(-123), val.Int16, "decoded int test")
		require.Equal(t, int64(456), val.Int64, "decoded int test")
		require.Equal(t, int32(-789), val.Int32, "decoded int test")
	})
}
