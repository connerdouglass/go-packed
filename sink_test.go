package packed_test

import (
	"testing"

	"github.com/connerdouglass/go-packed"
	"github.com/stretchr/testify/require"
)

func TestPackedSink(t *testing.T) {
	t.Run("sink", func(t *testing.T) {
		type testType struct {
			Str1 string
			Str2 string
			Str3 string
		}
		val := testType{
			Str1: "hello",
			Str2: "world",
			Str3: "!",
		}
		p := packed.Sequence(
			packed.String(&val.Str1),
			packed.String(&val.Str2),
			packed.String(&val.Str3),
		)
		encoded, err := p.Encode()
		require.NoError(t, err, "encoding string test")
		require.NotNil(t, encoded, "encoded string test")

		val.Str1 = ""
		val.Str2 = ""
		val.Str3 = ""

		p = packed.Sequence(
			packed.String(&val.Str1),
			packed.String(packed.Sink[string]()),
			packed.String(&val.Str3),
		)

		n, err := p.Decode(encoded)
		require.NoError(t, err, "decoding string test")
		require.Equal(t, len(encoded), n, "decoded string test")

		require.Equal(t, "hello", val.Str1, "decoded string test")
		require.Equal(t, "", val.Str2, "decoded string test")
		require.Equal(t, "!", val.Str3, "decoded string test")
	})
}
