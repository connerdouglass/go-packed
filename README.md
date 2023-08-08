# go-packed

Go library for packing and unpacking structured data to a binary format.

## Installation

```bash
go get github.com/connerdouglass/go-packed
```

## Examples

For these examples, assume the following type:

```go
type User struct {
    FirstName   string
    LastName    string
    Age         uint8
    BirthYear   uint16
    SocialLinks []string
}
```

### Encode

```go
user := User{
    FirstName:   "John",
    LastName:    "Doe",
    Age:         32,
    BirthYear:   1991,
    SocialLinks: []string{
        "https://facebook.com/johndoe",
        "https://twitter.com/johndoe",
    },
}

// Pack the user data to a byte slice
p := packed.Sequence(
    packed.String(&user.FirstName),
    packed.String(&user.LastName),
    packed.Uint8(&user.Age),
    packed.Uint16(&user.BirthYear),
    packed.Slice(packed.String)(&user.SocialLinks),
)
encoded, err := p.Encode()
```

### Decode

```go
var decodedUser User

// Unpack the encoded data to the user struct
p := packed.Sequence(
    packed.String(&decodedUser.FirstName),
    packed.String(&decodedUser.LastName),
    packed.Uint8(&decodedUser.Age),
    packed.Uint16(&decodedUser.BirthYear),
    packed.Slice(packed.String)(&decodedUser.SocialLinks),
)
_, err := p.Decode(encoded)
```

### Recommended Usage

For types you're frequently packing and unpacking (such as messages sent over the network), consider adding a method to those types (ie. `Packed()`) that returns a `packed.Packed` instance for that type.

This gives you a declarative way to define how your types should be packed and unpacked, so you can easily encode and decode those types.

```go
type ExampleMessage struct {
    MessageID   uint64
    SenderID    uint64
    MessageType uint8
    Content     string
    CreatedAt   time.Time
    ExpiresAt   time.Time
}

func (m *ExampleMessage) Packed() packed.Packed {
    return packed.Sequence(
        packed.Uint64(&m.MessageID),
        packed.Uint64(&m.SenderID),
        packed.Uint8(&m.MessageType),
        packed.String(&m.Content),
        packed.Time(&m.CreatedAt),
        packed.Time(&m.ExpiresAt),
    )
}
```

Then, you can easily encode an `ExampleMessage` object to bytes:

```go
message := ExampleMessage{...}
encodedBytes, err := message.Packed().Encode()
```

Or, you can decode bytes into an `ExampleMessage` object:

```go
var message ExampleMessage
_, err := message.Packed().Decode(encodedBytes)
```

## Nesting Packed Types

You can encode and decode structs that contain other structs.

```go
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
```

## Custom Types

You can extend the packed format with custom types. You just need to define an encode and decode function for the type.

```go
func PackMyCustomType(ptr *MyCustomType) packed.Packed {
	return packed.NewEncodeDecode(
		// Encode
		func() ([]byte, error) { ... },
		// Decode
		func(buf []byte) (int, error) { ... },
	)
}
```

Then you can use your custom type in a packed sequence:

```go
packed.Sequence(
    // ...
    PackMyCustomType(&myCustomType),
    // ...
)
```

## Why use this?

This packed format is smaller and faster to encode and decode than JSON. If you're sending messages rapidly over the network, such as in a multiplayer video game or other realtime application, it can be useful to use a packed format to reduce latency.

There's a very simple benchmark program in `cmd/benchmark` that compares the performance of JSON to packed. In that example, the packed format is 25x faster and 4x smaller than JSON.

```

## Why shouldn't you use this?

There are drawbacks. When transmitting packed data, both the sender and recipient need to be aware of the structure of the data. There's no metadata embedded in the encoded data. If the sender and recipient don't agree on the format, the recipient will not be able to make sense of the data.
