package util

// Ptr returns a pointer to the given value. Allows succinctly creating pointers to constants
func Ptr[T any](val T) *T {
	return &val
}
