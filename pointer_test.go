package packed_test

import (
	"log"
	"testing"

	"github.com/connerdouglass/go-packed"
	"github.com/stretchr/testify/require"
)

func ptr[T any](t T) *T {
	return &t
}

func TestPackedPointers(t *testing.T) {
	t.Run("pointers", func(t *testing.T) {
		type testType struct {
			Name     *string
			Age      *uint8
			EmptyStr *string
			NilStr   *string
		}
		val := testType{
			Name:     ptr("John"),
			Age:      ptr(uint8(30)),
			EmptyStr: ptr(""),
			NilStr:   nil,
		}
		encoded, err := packed.Sequence(
			packed.Pointer(packed.String)(&val.Name),
			packed.Pointer(packed.Uint8)(&val.Age),
			packed.Pointer(packed.String)(&val.EmptyStr),
			packed.Pointer(packed.String)(&val.NilStr),
		).Encode()
		require.NoError(t, err, "encoding pointer test")
		require.NotNil(t, encoded, "encoded pointer test")

		var dest testType
		n, err := packed.Sequence(
			packed.Pointer(packed.String)(&dest.Name),
			packed.Pointer(packed.Uint8)(&dest.Age),
			packed.Pointer(packed.String)(&dest.EmptyStr),
			packed.Pointer(packed.String)(&dest.NilStr),
		).Decode(encoded)
		require.NoError(t, err, "decoding pointer test")
		require.Equal(t, len(encoded), n, "decoded pointer test")

		log.Println("dest.Name", dest.Name)

		require.Equal(t, "John", *dest.Name, "decoded pointer test")
		require.Equal(t, uint8(30), *dest.Age, "decoded pointer test")
		require.Equal(t, "", *dest.EmptyStr, "decoded pointer test")
		require.Nil(t, dest.NilStr, "decoded pointer test")
	})
}
