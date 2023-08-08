package packed_test

import (
	"testing"
	"time"

	"github.com/connerdouglass/go-packed"
	"github.com/stretchr/testify/require"
)

func TestPackedTime(t *testing.T) {
	t.Run("time", func(t *testing.T) {
		type testType struct {
			Time1 time.Time
			Time2 time.Time
		}

		// Time encoding is accurate to the second, so we create a time without nanoseconds.
		time1 := time.Unix(time.Now().Unix(), 0)
		time2 := time1.Add(time.Minute * 137)
		val := testType{
			Time1: time1,
			Time2: time2,
		}
		p := packed.Sequence(
			packed.Time(&val.Time1),
			packed.Time(&val.Time2),
		)
		encoded, err := p.Encode()
		require.NoError(t, err, "encoding time test")
		require.NotNil(t, encoded, "encoded time test")

		n, err := p.Decode(encoded)
		require.NoError(t, err, "decoding time test")
		require.Equal(t, len(encoded), n, "decoded time test")

		require.Equal(t, time1, val.Time1, "decoded time test")
		require.Equal(t, time2, val.Time2, "decoded time test")
	})
}
