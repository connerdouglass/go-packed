package packed_test

import (
	"testing"

	"github.com/connerdouglass/go-packed"
	"github.com/stretchr/testify/require"
)

func TestPackedFloat(t *testing.T) {
	t.Run("float32", func(t *testing.T) {
		input1 := float32(123.456)
		input2 := float32(0.00123)
		encoded, err := packed.Sequence(
			packed.Float32(&input1),
			packed.Float32(&input2),
		).Encode()

		require.NoError(t, err, "float32 encode returned error")
		require.NotNil(t, encoded, "float32 encode returned nil")

		var output1, output2 float32

		n, err := packed.Sequence(
			packed.Float32(&output1),
			packed.Float32(&output2),
		).Decode(encoded)
		require.NoError(t, err, "float32 decode returned error")
		require.Equal(t, len(encoded), n, "float32 decode returned invalid length")

		require.Equal(t, float32(123.456), output1, "float32 output 1")
		require.Equal(t, float32(0.00123), output2, "float32 output 2")
	})
	t.Run("float64", func(t *testing.T) {
		input1 := float64(123.456)
		input2 := float64(0.00123)
		encoded, err := packed.Sequence(
			packed.Float64(&input1),
			packed.Float64(&input2),
		).Encode()

		require.NoError(t, err, "float64 encode returned error")
		require.NotNil(t, encoded, "float64 encode returned nil")

		var output1, output2 float64

		n, err := packed.Sequence(
			packed.Float64(&output1),
			packed.Float64(&output2),
		).Decode(encoded)
		require.NoError(t, err, "float64 decode returned error")
		require.Equal(t, len(encoded), n, "float64 decode returned invalid length")

		require.Equal(t, float64(123.456), output1, "float64 output 1")
		require.Equal(t, float64(0.00123), output2, "float64 output 2")
	})
}
