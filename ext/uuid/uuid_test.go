package uuid_test

import (
	"testing"

	"github.com/connerdouglass/go-packed"
	packeduuid "github.com/connerdouglass/go-packed/ext/uuid"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestPackedUUID(t *testing.T) {
	t.Run("uuid", func(t *testing.T) {
		type testType struct {
			MyUUID1 uuid.UUID
			MyUUID2 uuid.UUID
		}
		uuid1 := uuid.New()
		uuid2 := uuid.New()
		val := testType{
			MyUUID1: uuid1,
			MyUUID2: uuid2,
		}
		p := packed.Sequence(
			packeduuid.UUID(&val.MyUUID1),
			packeduuid.UUID(&val.MyUUID2),
		)
		encoded, err := p.Encode()
		require.NoError(t, err, "encoding uuid test")
		require.NotNil(t, encoded, "encoded uuid test")

		n, err := p.Decode(encoded)
		require.NoError(t, err, "decoding uuid test")
		require.Equal(t, len(encoded), n, "decoded uuid test")

		require.Equal(t, uuid1, val.MyUUID1, "decoded uuid test")
		require.Equal(t, uuid2, val.MyUUID2, "decoded uuid test")
	})
}
