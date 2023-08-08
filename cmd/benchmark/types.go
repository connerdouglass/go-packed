package main

import "github.com/connerdouglass/go-packed"

type Rectangle struct {
	ID       string   `json:"id"`
	Position Position `json:"position"`
	Size     Size     `json:"size"`
	Color    Color    `json:"color"`
}

type Position struct {
	X, Y, Z float32
}

type Size struct {
	Width, Height float32
}

type Color struct {
	R, G, B, A uint8
}

func (r *Rectangle) Packed() packed.Packed {
	return packed.Sequence(
		packed.String(&r.ID),
		r.Position.Packed(),
		r.Size.Packed(),
		r.Color.Packed(),
	)
}

func (s *Position) Packed() packed.Packed {
	return packed.Sequence(
		packed.Float32(&s.X),
		packed.Float32(&s.Y),
		packed.Float32(&s.Z),
	)
}

func (s *Size) Packed() packed.Packed {
	return packed.Sequence(
		packed.Float32(&s.Width),
		packed.Float32(&s.Height),
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
