package p_interface

type Addable[T any] interface {
	Add(other T) T
}
