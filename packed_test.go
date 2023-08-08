package packed_test

import (
	"testing"

	"github.com/connerdouglass/go-packed"
	"github.com/stretchr/testify/require"
)

func TestPacked(t *testing.T) {
	t.Run("structs", func(t *testing.T) {
		type testType struct {
			Name           string
			Age            uint8
			FavoriteNumber uint16
			Adult          bool
		}
		source := testType{
			Name:           "John Doe",
			Age:            42,
			FavoriteNumber: 1337,
			Adult:          true,
		}

		// Encode the source
		encoded, err := packed.Sequence(
			packed.String(&source.Name),
			packed.Uint8(&source.Age),
			packed.Uint16(&source.FavoriteNumber),
			packed.Bool(&source.Adult),
		).Encode()
		require.NoError(t, err, "encoding struct test")
		require.NotNil(t, encoded, "encoded struct test")

		// Decode into a destination
		var dest testType
		n, err := packed.Sequence(
			packed.String(&dest.Name),
			packed.Uint8(&dest.Age),
			packed.Uint16(&dest.FavoriteNumber),
			packed.Bool(&dest.Adult),
		).Decode(encoded)
		require.NoError(t, err, "decoding struct test")
		require.Equal(t, len(encoded), n, "decoded struct test")

		require.Equal(t, "John Doe", dest.Name, "decoded struct test")
		require.Equal(t, uint8(42), dest.Age, "decoded struct test")
		require.Equal(t, uint16(1337), dest.FavoriteNumber, "decoded struct test")
		require.Equal(t, true, dest.Adult, "decoded struct test")
	})
	t.Run("readme tests", func(t *testing.T) {
		rect := Rectangle{
			Size: Size{
				Width:  100,
				Height: 200,
			},
			Color: Color{
				R: 255,
				G: 127,
				B: 0,
				A: 255,
			},
		}
		encoded, err := rect.Packed().Encode()
		require.NoError(t, err, "encoding rectangle")
		require.NotNil(t, encoded, "encoded rectangle")

		var decoded Rectangle
		n, err := decoded.Packed().Decode(encoded)
		require.NoError(t, err, "decoding rectangle")
		require.Equal(t, len(encoded), n, "decoded rectangle")

		require.Equal(t, int32(100), decoded.Size.Width, "decoded rectangle")
		require.Equal(t, int32(200), decoded.Size.Height, "decoded rectangle")
		require.Equal(t, uint8(255), decoded.Color.R, "decoded rectangle")
		require.Equal(t, uint8(127), decoded.Color.G, "decoded rectangle")
		require.Equal(t, uint8(0), decoded.Color.B, "decoded rectangle")
		require.Equal(t, uint8(255), decoded.Color.A, "decoded rectangle")
	})
	t.Run("readme tests with pointers", func(t *testing.T) {
		rect := RectangleWithPointers{
			Size: &Size{
				Width:  100,
				Height: 200,
			},
			Color: nil,
		}
		encoded, err := rect.Packed().Encode()
		require.NoError(t, err, "encoding rectangle with pointers")
		require.NotNil(t, encoded, "encoded rectangle with pointers")

		var decoded RectangleWithPointers
		n, err := decoded.Packed().Decode(encoded)
		require.NoError(t, err, "decoding rectangle with pointers")
		require.Equal(t, len(encoded), n, "decoded rectangle with pointers")

		require.Equal(t, int32(100), decoded.Size.Width, "decoded rectangle with pointers")
		require.Equal(t, int32(200), decoded.Size.Height, "decoded rectangle with pointers")
		require.Nil(t, decoded.Color, "decoded rectangle with pointers")
	})
}

type Rectangle struct {
	Size  Size
	Color Color
}

type Size struct {
	Width, Height int32
}

type Color struct {
	R, G, B, A uint8
}

func (r *Rectangle) Packed() packed.Packed {
	return packed.Sequence(
		r.Size.Packed(),
		r.Color.Packed(),
	)
}

func (s *Size) Packed() packed.Packed {
	return packed.Sequence(
		packed.Int32(&s.Width),
		packed.Int32(&s.Height),
	)
}

func (c *Color) Packed() packed.Packed {
	return packed.Sequence(
		packed.Uint8(&c.R),
		packed.Uint8(&c.G),
		packed.Uint8(&c.B),
		packed.Uint8(&c.A),
	)
}

type RectangleWithPointers struct {
	Size  *Size
	Color *Color
}

func (r *RectangleWithPointers) Packed() packed.Packed {
	return packed.Sequence(
		packed.Pointer((*Size).Packed)(&r.Size),
		packed.Pointer((*Color).Packed)(&r.Color),
	)
}
