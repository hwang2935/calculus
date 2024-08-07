package p_interface

type Containable[T any] interface {
	Contains(value T) bool
}
