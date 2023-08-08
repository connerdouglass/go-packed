package packed_test

import (
	"testing"

	"github.com/connerdouglass/go-packed"
	"github.com/stretchr/testify/require"
)

func TestPackedBool(t *testing.T) {
	t.Run("bool", func(t *testing.T) {
		input1 := false
		input2 := true
		encoded, err := packed.Sequence(
			packed.Bool(&input1),
			packed.Bool(&input2),
		).Encode()

		require.NoError(t, err, "bool encode returned error")
		require.NotNil(t, encoded, "bool encode returned nil")

		var output1, output2 bool

		n, err := packed.Sequence(
			packed.Bool(&output1),
			packed.Bool(&output2),
		).Decode(encoded)
		require.NoError(t, err, "bool decode returned error")
		require.Equal(t, len(encoded), n, "bool decode returned invalid length")

		require.Equal(t, false, output1, "bool output 1")
		require.Equal(t, true, output2, "bool output 2")
	})
}
