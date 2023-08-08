package packed_test

import (
	"testing"

	"github.com/connerdouglass/go-packed"
	"github.com/stretchr/testify/require"
)

func TestPackedMap(t *testing.T) {
	t.Run("maps", func(t *testing.T) {

		inputMap := map[string]uint8{
			"hello": 1,
			"world": 2,
			"foo":   3,
			"bar":   4,
		}

		encoded, err := packed.Map(packed.String, packed.Uint8)(&inputMap).Encode()
		require.NoError(t, err, "encoding maps test")
		require.NotNil(t, encoded, "encoded maps test")

		var outputMap map[string]uint8

		n, err := packed.Map(packed.String, packed.Uint8)(&outputMap).Decode(encoded)
		require.NoError(t, err, "decoding maps test")
		require.Equal(t, len(encoded), n, "decoded maps test")

		require.Equal(t, 4, len(outputMap), "decoded maps test")
		require.Equal(t, uint8(1), outputMap["hello"], "decoded maps test")
		require.Equal(t, uint8(2), outputMap["world"], "decoded maps test")
		require.Equal(t, uint8(3), outputMap["foo"], "decoded maps test")
		require.Equal(t, uint8(4), outputMap["bar"], "decoded maps test")
	})
}
