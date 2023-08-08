package packed

// Sink is a helper function for creating a sink for a type. This allows you to decode
// but not retain a value that has been removed from the struct, but has not been removed
// from the encoded data. Useful for backwards compatibility when removing fields.
func Sink[T any]() *T {
	var val T
	return &val
}
