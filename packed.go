package packed

// Packed is a value that can be encoded and decoded.
type Packed interface {
	Encode() ([]byte, error)
	Decode([]byte) (int, error)
}

// Type is a function that can bind to a pointer and encode or decode it.
type Type[T any] func(*T) Packed

// Sequence creates a Packed that can encode and decode sequences of values.
func Sequence(fields ...Packed) Packed {
	return &packedSequence{fields}
}

type packedSequence struct {
	fields []Packed
}

func (p *packedSequence) Encode() ([]byte, error) {
	var buf []byte
	for _, field := range p.fields {
		fieldEncoded, err := field.Encode()
		if err != nil {
			return nil, err
		}
		buf = append(buf, fieldEncoded...)
	}
	return buf, nil
}

func (p *packedSequence) Decode(buf []byte) (int, error) {
	var cursor int
	for _, field := range p.fields {
		decoded, err := field.Decode(buf[cursor:])
		if err != nil {
			return 0, err
		}
		cursor += decoded
	}
	return cursor, nil
}

// NewEncodeDecode creates a Packed that can encode and decode values using the provided functions.
func NewEncodeDecode(
	encode func() ([]byte, error),
	decode func([]byte) (int, error),
) Packed {
	return &packedFuncs{encode, decode}
}

type packedFuncs struct {
	encode func() ([]byte, error)
	decode func([]byte) (int, error)
}

func (p *packedFuncs) Encode() ([]byte, error) {
	return p.encode()
}

func (p *packedFuncs) Decode(buf []byte) (int, error) {
	return p.decode(buf)
}
